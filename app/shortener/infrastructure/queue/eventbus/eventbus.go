package eventbus

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/application/event"
	"github.com/creepzed/url-shortener-service/app/shared/domain/queue"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
)

type eventBusKafka struct {
	queue   queue.PublisherQueue
	topic   string
	channel chan *queue.MessageBase
}

func NewEventBusKafka(publisher queue.PublisherQueue, topic string) *eventBusKafka {
	return &eventBusKafka{
		queue:   publisher,
		topic:   topic,
		channel: make(chan *queue.MessageBase, 100),
	}
}

func (e *eventBusKafka) Publish(ctx context.Context, events []event.Event) error {
	for _, v := range events {
		messageData, err := queue.NewMessageData(v.AggregateID(), v)
		if err != nil {
			return err
		}
		e.channel <- messageData
	}
	return nil
}

func (e eventBusKafka) Observer(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			log.Debug("shutting down observer")
			return

		case messageData := <-e.channel:
			e.queue.Publish(ctx, e.topic, messageData)
		}
	}
}
