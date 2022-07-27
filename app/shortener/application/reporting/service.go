package reporting

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/application/query"
	"github.com/creepzed/url-shortener-service/app/shared/domain/queue"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
	"github.com/creepzed/url-shortener-service/app/shortener/application/finding"
	"github.com/google/uuid"
)

type MessageStatics struct {
	UrlId    string                 `json:"url_id"`
	Result   interface{}            `json:"result"`
	Metadata map[string]interface{} `json:"metadata"`
}

type reportWrapApplicationService struct {
	service  finding.FindApplicationService
	producer queue.PublisherQueue
	topic    string
	channel  chan *MessageStatics
}

func NewReportApplicationService(service finding.FindApplicationService, producer queue.PublisherQueue, topic string) *reportWrapApplicationService {
	return &reportWrapApplicationService{
		service:  service,
		producer: producer,
		topic:    topic,
		channel:  make(chan *MessageStatics, 100),
	}
}

func (ras reportWrapApplicationService) Do(ctx context.Context, query finding.FindUrlShortenerQuery) (query.Result, error) {

	result, err := ras.service.Do(ctx, query)
	if err == nil {
		ras.Publish(query, result)
	}
	
	return result, err
}

func (ras reportWrapApplicationService) Publish(query finding.FindUrlShortenerQuery, result query.Result) {
	statics := MessageStatics{
		UrlId:    query.UrlId(),
		Result:   result,
		Metadata: query.Metadata(),
	}

	ras.channel <- &statics
}

func (ras reportWrapApplicationService) Observer(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Debug("shutting down observer")
			return

		case statics := <-ras.channel:
			ras.PushMessage(ctx, statics)
		}
	}
}

func (ras reportWrapApplicationService) PushMessage(ctx context.Context, statics *MessageStatics) {
	messageData, errMessage := queue.NewMessageData(uuid.NewString(), statics)
	if errMessage != nil {
		log.WithError(errMessage).Error("error creating message")
	}
	errQueue := ras.producer.Publish(ctx, ras.topic, messageData)
	if errQueue != nil {
		log.WithError(errQueue).Error("error publishing message")
	}
}
