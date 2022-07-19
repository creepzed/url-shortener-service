package inmemory

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/application/event"
)

type eventBusInMemory struct {
	queue stack
}

func NewEventBusInMemory() *eventBusInMemory {
	return &eventBusInMemory{
		queue: stack{},
	}
}

func (e *eventBusInMemory) Publish(ctx context.Context, events []event.Event) error {
	e.queue = events
	return nil
}

type stack []event.Event

func (s *stack) IsEmpty() bool {
	return len(*s) == 0
}

func (s *stack) Push(str event.Event) {
	*s = append(*s, str)
}

func (s *stack) Pop() (*event.Event, bool) {
	if s.IsEmpty() {
		return nil, false
	} else {
		index := len(*s) - 1
		element := (*s)[index]
		*s = (*s)[:index]
		return &element, true
	}
}
