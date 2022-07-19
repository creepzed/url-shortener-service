package inmemory

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/application/command"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
)

type CommandBus struct {
	handlers map[command.Type]command.Handler
}

func NewCommandBusMemory() *CommandBus {
	return &CommandBus{
		handlers: make(map[command.Type]command.Handler),
	}
}

func (cb *CommandBus) Dispatch(ctx context.Context, cmd command.Command) error {
	handler, ok := cb.handlers[cmd.Type()]
	if !ok {
		return nil
	}
	err := handler.Handle(ctx, cmd)
	if err != nil {
		log.Error("error while handling %s - %s\n", cmd.Type(), err)
		return err
	}
	return nil
}

func (cb CommandBus) Register(cmdType command.Type, handler command.Handler) {
	cb.handlers[cmdType] = handler
}
