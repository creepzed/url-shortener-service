package eventbus

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/application/event"
	"github.com/creepzed/url-shortener-service/app/shared/domain/queue"
)

type eventBusKafka struct {
	queue queue.PublisherQueue
	topic string
}

func NewEventBusKafka(queue queue.PublisherQueue, topic string) *eventBusKafka {
	return &eventBusKafka{
		queue: queue,
		topic: topic,
	}
}

func (e *eventBusKafka) Publish(ctx context.Context, events []event.Event) error {
	for _, v := range events {
		messageData, err := queue.NewMessageData(v.AggregateID(), v)
		if err != nil {
			return err
		}
		e.queue.Publish(ctx, e.topic, messageData)
	}
	return nil
}