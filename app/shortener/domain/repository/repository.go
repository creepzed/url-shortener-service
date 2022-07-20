package repository

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
)

type KeyGenerateService interface {
	GetKey() (vo.UrlId, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=../../infrastructure/storage/storagemocks --name=KeyGenerateService

type UrlShortenerRepository interface {
	Create(ctx context.Context, urlShortener domain.UrlShortener) error
	FindById(ctx context.Context, urlId vo.UrlId) (domain.UrlShortener, error)
}

//go:generate mockery --case=snake --outpkg=storagemocks --output=../../infrastructure/storage/storagemocks --name=UrlShortenerRepository
