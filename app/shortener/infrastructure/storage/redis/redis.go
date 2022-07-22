package redis

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/storage/redisdb"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/utils"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
	"time"
)

type UrlShortenerRedis struct {
	UrlId       string `json:"url_id"`
	UrlEnable   bool   `json:"url_enable"`
	OriginalUrl string `json:"original_url"`
	UserId      string `json:"user_id"`
}

func NewUrlShortenerRedis(urlShortener domain.UrlShortener) *UrlShortenerRedis {
	return &UrlShortenerRedis{
		UrlId:       urlShortener.UrlId().Value(),
		UrlEnable:   urlShortener.IsEnabled().Value(),
		OriginalUrl: urlShortener.OriginalUrl().Value(),
		UserId:      urlShortener.UrlId().Value(),
	}
}

type urlShortenerRepositoryRedis struct {
	baseRepository *redisdb.RepositoryRedis
}

func NewUrlShortenerRepositoryRedis(connection redisdb.ConnectionRedis, dbTimeout time.Duration) *urlShortenerRepositoryRedis {
	return &urlShortenerRepositoryRedis{
		baseRepository: redisdb.NewRepositoryRedis(connection, dbTimeout),
	}
}

func (r *urlShortenerRepositoryRedis) Create(ctx context.Context, urlShortener domain.UrlShortener) error {
	doc := NewUrlShortenerRedis(urlShortener)

	err := r.baseRepository.Set(ctx, urlShortener.UrlId().Value(), doc)
	if err != nil {
		return err
	}
	return nil
}

func (r *urlShortenerRepositoryRedis) FindById(ctx context.Context, urlId vo.UrlId) (domain.UrlShortener, error) {
	result, err := r.baseRepository.Get(ctx, urlId.Value())
	if err != nil {
		return domain.UrlShortener{}, err
	}

	doc := new(UrlShortenerRedis)
	err = utils.ConvertEntity(result, &doc)
	if err != nil {
		return domain.UrlShortener{}, err
	}
	anUrlId, err := vo.NewUrlId(doc.UrlId)
	if err != nil {
		return domain.UrlShortener{}, err
	}

	anUrlEnable := vo.NewUrlEnabled(doc.UrlEnable)

	anOriginalUrl, err := vo.NewOriginalUrl(doc.OriginalUrl)
	if err != nil {
		return domain.UrlShortener{}, err
	}

	anUserId, err := vo.NewUserId(doc.UserId)
	if err != nil {
		return domain.UrlShortener{}, err
	}

	anUrlShortener := domain.LoadUrlShortener(anUrlId, anUrlEnable, anOriginalUrl, anUserId)
	return anUrlShortener, nil
}
