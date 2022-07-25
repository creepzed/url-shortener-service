package subscriber

import (
	"context"
	"encoding/json"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/storage/redisdb"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/storage/redis"
	"github.com/segmentio/kafka-go"
	"time"
)

type message struct {
	EventId     string    `json:"eventId"`
	EventType   string    `json:"event_type"`
	AggregateId string    `json:"aggregate_id"`
	OccurredOn  time.Time `json:"occurred_on"`
	UrlId       string    `json:"url_id"`
	UrlStatus   bool      `json:"url_status"`
	OriginUrl   string    `json:"origin_url"`
	UserId      string    `json:"user_id"`
}

type subscriberUpdateCache struct {
	dialer  *kafka.Dialer
	groupID string
	brokers []string
	topic   string
	cache   *redisdb.RepositoryRedis
}

func NewSubscriberUpdateCache(cache *redisdb.RepositoryRedis, groupID string, topic string, dialer *kafka.Dialer, brokers ...string) *subscriberUpdateCache {
	return &subscriberUpdateCache{
		dialer:  dialer,
		groupID: groupID,
		brokers: brokers,
		topic:   topic,
		cache:   cache,
	}
}

func (s *subscriberUpdateCache) getKafkaReader(topic string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:        s.brokers,
		GroupID:        s.groupID,
		Topic:          topic,
		MinBytes:       1e6,  // 1MB
		MaxBytes:       10e6, // 10MB
		CommitInterval: time.Second,
		Dialer:         s.dialer,
	})
}

func (s *subscriberUpdateCache) ReadMessage(ctx context.Context) {
	reader := s.getKafkaReader(s.topic)
	defer reader.Close()
	for {
		select {
		case <-ctx.Done():
			log.Debug("shutting down subscriber")
			return
		default:
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				log.WithError(err).Error("error reading kafka messages")
				continue
			}
			payload := new(message)
			err = json.Unmarshal(msg.Value, &payload)
			if err != nil {
				log.WithError(err).Error("error unmarshalling kafka messages")
				continue
			}

			doc := redis.UrlShortenerRedis{
				UrlId:       payload.UrlId,
				UrlEnable:   payload.UrlStatus,
				OriginalUrl: payload.OriginUrl,
				UserId:      payload.UserId,
			}

			err = s.cache.Set(ctx, doc.UrlId, doc)
			if err != nil {
				log.WithError(err).Error("error saving on cache")
				continue
			}
		}
	}
}
