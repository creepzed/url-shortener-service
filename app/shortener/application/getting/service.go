package getting

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/application/query"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/repository"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/transformer"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
)

type GetAllApplicationService interface {
	Do(ctx context.Context, query GetAllUrlShortenerQuery) (query.Result, error)
}

//go:generate mockery --case=snake --outpkg=servicemocks --output=../mocks/servicemocks --name=GetAllApplicationService

type getAllApplicationService struct {
	repository  repository.GetAllByUserIdRepository
	transformer transformer.UrlShortenerTransformer
}

func NewGetAllApplicationService(repository repository.GetAllByUserIdRepository, transformer transformer.UrlShortenerTransformer) *getAllApplicationService {
	return &getAllApplicationService{
		repository:  repository,
		transformer: transformer,
	}
}

func (fas getAllApplicationService) Do(ctx context.Context, query GetAllUrlShortenerQuery) (query.Result, error) {
	userId, err := vo.NewUserId(query.UserId())
	if err != nil {
		return nil, err
	}

	urlShortList, err := fas.repository.GetAllByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}

	result, err := fas.transformer.TransformList(urlShortList)
	if err != nil {
		return nil, err
	}

	return result, nil
}
