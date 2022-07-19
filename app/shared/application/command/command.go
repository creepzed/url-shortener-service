package command

import "context"

type Bus interface {
	// Dispatch is the method used to dispatch new commands.
	Dispatch(context.Context, Command) error
	// Register is the method used to register a new command handler.
	Register(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=commandmocks --output=../../infrastructure/bus/busMocks/commandmocks --name=Bus

type Type string

type Command interface {
	Type() Type
}

type Handler interface {
	Handle(context.Context, Command) error
}

type CommandBus interface {
	Dispatch(ctx context.Context, cmd Command) error
	Register(cmdType Type, handler Handler)
}
