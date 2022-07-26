package getting

import (
	"context"
	"errors"
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/application/query"
)

type GetAllUrlShortenerQueryHandler struct {
	applicationService GetAllApplicationService
}

var (
	ErrUnexpectedQuery = errors.New("unexpected query")
)

func NewGetAllUrlShortenerQueryHandler(applicationService GetAllApplicationService) GetAllUrlShortenerQueryHandler {
	return GetAllUrlShortenerQueryHandler{
		applicationService: applicationService,
	}
}

func (h GetAllUrlShortenerQueryHandler) Handle(ctx context.Context, qry query.Query) (query.Result, error) {
	query, ok := qry.(GetAllUrlShortenerQuery)
	if !ok {
		return nil, fmt.Errorf("%w: %s", ErrUnexpectedQuery, qry.Type())
	}
	return h.applicationService.Do(ctx, query)
}
