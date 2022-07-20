package finding

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/application/query"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/repository"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/transformer"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
)

type FindApplicationService interface {
	Do(ctx context.Context, query FindUrlShortenerQuery) (query.Result, error)
}

type findApplicationService struct {
	repository  repository.UrlShortenerRepository
	transformer transformer.UrlShortenerTransformer
}

func NewFindApplicationService(repository repository.UrlShortenerRepository, transformer transformer.UrlShortenerTransformer) *findApplicationService {
	return &findApplicationService{
		repository:  repository,
		transformer: transformer,
	}
}

func (fas findApplicationService) Do(ctx context.Context, query FindUrlShortenerQuery) (query.Result, error) {
	urlId, err := vo.NewUrlId(query.UrlId())
	if err != nil {
		return nil, err
	}

	urlShortener, err := fas.repository.FindById(ctx, urlId)
	if err != nil {
		return domain.UrlShortener{}, err
	}

	resultUrl, err := fas.transformer.Transform(urlShortener)
	if err != nil {
		return domain.UrlShortener{}, err
	}
	
	return resultUrl, nil
}
