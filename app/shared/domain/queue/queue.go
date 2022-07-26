package queue

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type MessageData interface {
	Key() []byte
	Payload() []byte
}

type MessageBase struct {
	key     []byte
	payload []byte
}

func NewMessageData(key string, payload interface{}) (*MessageBase, error) {

	if len(key) == 0 {
		return nil, errors.New("error key is empty")
	}

	payloadString, errConvertPayload := json.Marshal(payload)
	if errConvertPayload != nil {
		return nil, fmt.Errorf("data marshaling error: %s", errConvertPayload.Error())
	}
	event := &MessageBase{
		key:     []byte(key),
		payload: payloadString,
	}
	return event, nil
}

func (e *MessageBase) Key() []byte {
	return e.key
}

func (e *MessageBase) Payload() []byte {
	return e.payload
}

type PublisherQueue interface {
	Publish(ctx context.Context, topic string, messageData MessageData) error
	Close(topic string) error
}

//go:generate mockery --case=snake --outpkg=queuemocks --output=../mocks/queuemocksmocks --name=PublisherQueue
