package cache

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/repository"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
)

type repositoryCached struct {
	db    repository.FindByIdUrlShortenerRepository
	cache repository.CreateAndFindRepository
}

func NewCache(repositoryDB repository.FindByIdUrlShortenerRepository,
	repositoryCache repository.CreateAndFindRepository) *repositoryCached {
	return &repositoryCached{
		db:    repositoryDB,
		cache: repositoryCache,
	}
}

func (r *repositoryCached) FindById(ctx context.Context, urlId vo.UrlId) (domain.UrlShortener, error) {
	resultCache, err := r.cache.FindById(ctx, urlId)
	if err != nil {
		resultDB, err := r.db.FindById(ctx, urlId)
		if err == nil {
			errCache := r.cache.Create(ctx, resultDB)
			if errCache != nil {
				log.WithError(err).Error("error saving in cache")
			}
		}
		return resultDB, err

	}
	return resultCache, nil

}
