package inmemory

import (
	"context"
	"errors"
	"github.com/creepzed/url-shortener-service/app/shared/application/query"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
)

type QueryBus struct {
	handlers map[query.Type]query.Handler
}

func NewQueryBusMemory() *QueryBus {
	return &QueryBus{
		handlers: make(map[query.Type]query.Handler),
	}
}

func (qb QueryBus) Execute(ctx context.Context, qry query.Query) (query.Result, error) {
	handler, ok := qb.handlers[qry.Type()]
	if !ok {
		return nil, errors.New("error: query not found")
	}
	result, err := handler.Handle(ctx, qry)
	if err != nil {
		log.Error("error: while query handling %s - %s", qry.Type(), err)
		return nil, err
	}
	return result, nil
}

func (qb QueryBus) Register(qryType query.Type, handler query.Handler) {
	qb.handlers[qryType] = handler
}
