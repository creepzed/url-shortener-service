package query

import "context"

type Type string

type Result interface{}

type Query interface {
	Type() Type
}

type Handler interface {
	Handle(context.Context, Query) (*Result, error)
}

type QueryBus interface {
	Execute(ctx context.Context, qry Query) (Result, error)
	Register(qryType Type, handler Handler)
}
