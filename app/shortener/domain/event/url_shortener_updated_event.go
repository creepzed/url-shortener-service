package event

import (
	"encoding/json"
	"github.com/creepzed/url-shortener-service/app/shared/application/event"
	"time"
)

const ShortenerUpdatedEventType event.Type = "event.url.shortener.updated"

type ShortenerUpdatedEvent struct {
	urlId       string
	urlStatus   bool
	originalUrl string
	userId      string

	BaseEvent
}

func NewShortenerUpdatedEvent(urlId string, urlStatus bool, originalUrl string, userId string) ShortenerUpdatedEvent {
	return ShortenerUpdatedEvent{
		urlId:       urlId,
		urlStatus:   urlStatus,
		originalUrl: originalUrl,
		userId:      userId,
		BaseEvent:   NewBaseEvent(urlId),
	}
}

func (sce ShortenerUpdatedEvent) Type() event.Type {
	return ShortenerUpdatedEventType
}

func (sce ShortenerUpdatedEvent) UrlId() string {
	return sce.urlId
}

func (sce ShortenerUpdatedEvent) UrlStatus() bool {
	return sce.urlStatus
}

func (sce ShortenerUpdatedEvent) OriginalUrl() string {
	return sce.originalUrl
}

func (sce ShortenerUpdatedEvent) UserId() string {
	return sce.userId
}

func (sce ShortenerUpdatedEvent) MarshalJSON() ([]byte, error) {
	j, err := json.Marshal(&struct {
		EventId     string    `json:"eventId"`
		EventType   string    `json:"event_type"`
		AggregateId string    `json:"aggregate_id"`
		OccurredOn  time.Time `json:"occurred_on"`
		UrlId       string    `json:"url_id"`
		UrlStatus   bool      `json:"url_status"`
		OriginalUrl string    `json:"origin_url"`
		UserId      string    `json:"user_id"`
	}{
		EventId:     sce.eventID,
		EventType:   string(sce.Type()),
		AggregateId: sce.aggregateID,
		OccurredOn:  sce.occurredOn,
		UrlId:       sce.urlId,
		UrlStatus:   sce.urlStatus,
		OriginalUrl: sce.originalUrl,
		UserId:      sce.userId,
	})

	if err != nil {
		return nil, err
	}

	return j, err
}

func (sce ShortenerUpdatedEvent) UnmarshalJSON(b []byte) error {
	var value struct {
		EventId     string    `json:"eventId"`
		EventType   string    `json:"event_type"`
		AggregateId string    `json:"aggregate_id"`
		OccurredOn  time.Time `json:"occurred_on"`
		UrlId       string    `json:"url_id"`
		UrlStatus   bool      `json:"url_status"`
		OriginalUrl string    `json:"origin_url"`
		UserId      string    `json:"user_id"`
	}

	err := json.Unmarshal(b, &value)
	if err != nil {
		return err
	}

	sce.eventID = value.EventId
	sce.aggregateID = value.AggregateId
	sce.occurredOn = value.OccurredOn

	sce.urlId = value.UrlId
	sce.urlStatus = value.UrlStatus
	sce.originalUrl = value.OriginalUrl
	sce.userId = value.UserId

	return nil
}
