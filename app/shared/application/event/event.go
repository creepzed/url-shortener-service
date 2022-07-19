package event

import (
	"context"
	"time"
)

type Bus interface {
	Publish(context.Context, []Event) error
	//Subscribe(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=eventmocks --output=../../infrastructure/bus/busMocks/eventmocks --name=Bus

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
