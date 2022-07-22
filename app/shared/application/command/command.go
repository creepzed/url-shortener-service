package command

import "context"

type CommandBus interface {
	// Dispatch is the method used to dispatch new commands.
	Dispatch(context.Context, Command) error
	// Register is the method used to register a new command handler.
	Register(Type, Handler)
}

//go:generate mockery --case=snake --outpkg=commandmocks --output=../mocks/commandmocks --name=CommandBus

type Type string

type Command interface {
	Type() Type
}

//go:generate mockery --case=snake --outpkg=commandmocks --output=../mocks/commandmocks --name=Command

type Handler interface {
	Handle(context.Context, Command) error
}
