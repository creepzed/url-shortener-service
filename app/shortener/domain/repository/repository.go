package repository

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	shortenerVo "github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
)

type UrlShortenerRepository interface {
	CreateUrlShortenerRepository
	FindByIdUrlShortenerRepository
	UpdateUrlShortenerRepository
}

type CreateAndFindRepository interface {
	CreateUrlShortenerRepository
	FindByIdUrlShortenerRepository
}

type UpdateAndFindRepository interface {
	UpdateUrlShortenerRepository
	FindByIdUrlShortenerRepository
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=../mocks/storagemocks --name=UrlShortenerRepository

type CreateUrlShortenerRepository interface {
	Create(ctx context.Context, urlShortener domain.UrlShortener) error
}

type FindByIdUrlShortenerRepository interface {
	FindById(ctx context.Context, urlId vo.UrlId) (domain.UrlShortener, error)
}

type GetAllByUserIdRepository interface {
	GetAllByUserId(ctx context.Context, userId shortenerVo.UserId) ([]domain.UrlShortener, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=../mocks/storagemocks --name=GetAllByUserIdRepository

type UpdateUrlShortenerRepository interface {
	Update(ctx context.Context, urlShortener domain.UrlShortener) error
}
