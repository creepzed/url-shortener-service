package finding

import (
	"context"
	"errors"
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/application/query"
)

type FindUrlShortenerQueryHandler struct {
	applicationService FindApplicationService
}

var (
	ErrUnexpectedQuery = errors.New("unexpected query")
)

func NewFindUrlShortenerQueryHandler(applicationService FindApplicationService) FindUrlShortenerQueryHandler {
	return FindUrlShortenerQueryHandler{
		applicationService: applicationService,
	}
}

func (h FindUrlShortenerQueryHandler) Handle(ctx context.Context, qry query.Query) (query.Result, error) {
	command, ok := qry.(FindUrlShortenerQuery)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrUnexpectedQuery, qry.Type())
	}
	return h.applicationService.Do(ctx, command)
}
