package event

import (
	"context"
	"time"
)

type EventBus interface {
	Publish(context.Context, []Event) error
	//Subscribe(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=eventmocks --output=../mocks/eventmocks --name=EventBus

type Handler interface {
	Handle(context.Context, Event) error
}

type Type string

type Event interface {
	EventId() string
	AggregateID() string
	OccurredOn() time.Time
	Type() Type
}
