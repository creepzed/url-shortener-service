package inmemory

import (
	"context"
	"errors"
	"github.com/creepzed/url-shortener-service/app/shared/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
)

type urlShortenerRepositoryInMemory struct {
	data map[string]domain.UrlShortener
}

func NewDataBaseInMemory() *urlShortenerRepositoryInMemory {
	return &urlShortenerRepositoryInMemory{
		data: make(map[string]domain.UrlShortener),
	}
}

func (d *urlShortenerRepositoryInMemory) Create(ctx context.Context, urlShortener domain.UrlShortener) error {
	if _, found := d.data[urlShortener.UrlId().Value()]; found {
		return errors.New("the key is duplicated")
	}

	d.data[urlShortener.UrlId().Value()] = urlShortener
	return nil
}

func (d *urlShortenerRepositoryInMemory) FindById(ctx context.Context, urlId vo.UrlId) (domain.UrlShortener, error) {
	if v, found := d.data[urlId.Value()]; found {
		return v, nil
	}
	return domain.UrlShortener{}, nil
}
