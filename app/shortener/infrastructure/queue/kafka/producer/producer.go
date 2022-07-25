package producer

import (
	"context"
	"errors"
	"github.com/creepzed/url-shortener-service/app/shared/domain/queue"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/utils"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/compress"
)

type publisher struct {
	brokers      []string
	dialer       *kafka.Dialer
	kafkaWriters map[string]*kafka.Writer
}

func NewKafkaPublisher(kafkaDialer *kafka.Dialer, brokers ...string) *publisher {
	return &publisher{
		brokers:      brokers,
		kafkaWriters: make(map[string]*kafka.Writer),
		dialer:       kafkaDialer,
	}
}

func (p *publisher) getKafkaWriter(topic string) *kafka.Writer {
	if p.kafkaWriters[topic] == nil {
		p.kafkaWriters[topic] = &kafka.Writer{
			Addr:     kafka.TCP(p.brokers...),
			Topic:    topic,
			Balancer: &kafka.Hash{},
			//Balancer:    &kafka.LeastBytes{},
			Compression: compress.Snappy,
			Logger:      log.Logger(),
			ErrorLogger: log.Logger(),
		}
	}
	return p.kafkaWriters[topic]
}

func (p *publisher) Publish(ctx context.Context, topic string, eventData queue.MessageData) error {

	kafkaMessages, err := createKafkaMessages(eventData)
	if err != nil {
		return errors.New("error on publishing kafka message")
	}

	kafkaWriter := p.getKafkaWriter(topic)

	err = kafkaWriter.WriteMessages(ctx, kafkaMessages...)
	if err != nil {
		log.WithError(err).Error("error writing kafka message")
		return err
	} else {
		log.WithField("topic", topic).Info("kafka message published successfully")
	}

	return nil
}

func createKafkaMessages(data ...queue.MessageData) ([]kafka.Message, error) {
	var kafkaMessages []kafka.Message

	if utils.IsNilFixed(data) {
		return nil, errors.New("kafka error the data is empty")
	}

	for _, v := range data {
		kafkaMessages = append(kafkaMessages, kafka.Message{
			Key:   v.Key(),
			Value: v.Payload(),
		})
	}
	return kafkaMessages, nil
}

func (p *publisher) Close(topic string) error {
	if p.kafkaWriters[topic] == nil {
		return errors.New("error trying to close kafka connection, connection does not exist")
	}
	p.kafkaWriters[topic].Close()
	delete(p.kafkaWriters, topic)
	return nil
}
