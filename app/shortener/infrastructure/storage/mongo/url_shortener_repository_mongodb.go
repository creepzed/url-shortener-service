package mongo

import (
	"context"
	"fmt"
	vo2 "github.com/creepzed/url-shortener-service/app/shared/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/storage"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/storage/mongodb"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/utils"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type UrlShortenerMongoDB struct {
	UrlId       string `json:"url_id" bson:"url_id"`
	UrlEnable   bool   `json:"url_enable" bson:"url_enable"`
	OriginalUrl string `json:"original_url" bson:"original_url"`
	UserId      string `json:"user_id" bson:"user_id"`
}

func NewUrlShortenerMongo(urlShortener domain.UrlShortener) *UrlShortenerMongoDB {
	return &UrlShortenerMongoDB{
		UrlId:       urlShortener.UrlId().Value(),
		UrlEnable:   urlShortener.IsEnabled().Value(),
		OriginalUrl: urlShortener.OriginalUrl().Value(),
		UserId:      urlShortener.UserId().Value(),
	}
}

type urlShortenerRepositoryMongoDB struct {
	baseRepository storage.Repository
}

func NewUrlShortenerRepositoryMongo(connection mongodb.ConnectionMongo, dbTimeout time.Duration) *urlShortenerRepositoryMongoDB {
	return &urlShortenerRepositoryMongoDB{
		baseRepository: mongodb.NewRepositoryMongoDB(connection, dbTimeout),
	}
}

func (u *urlShortenerRepositoryMongoDB) Create(ctx context.Context, urlShortener domain.UrlShortener) error {
	doc := NewUrlShortenerMongo(urlShortener)

	errInsert := u.baseRepository.Create(ctx, doc)
	if errInsert != nil {
		if mongo.IsDuplicateKeyError(errInsert) {
			return fmt.Errorf("%w: %s", exception.ErrUrlIdDuplicate, urlShortener.UrlId())
		}
		return errInsert
	}
	return nil
}

func (u *urlShortenerRepositoryMongoDB) FindById(ctx context.Context, urlId vo2.UrlId) (domain.UrlShortener, error) {
	filter := map[string]interface{}{
		"url_id": urlId.Value(),
	}

	result, err := u.baseRepository.FindById(ctx, filter)
	if err != nil {
		return domain.UrlShortener{}, err
	}

	doc := new(UrlShortenerMongoDB)
	err = utils.ConvertEntity(result, &doc)
	if err != nil {
		return domain.UrlShortener{}, err
	}

	anUrlId, err := vo2.NewUrlId(doc.UrlId)
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

func (u *urlShortenerRepositoryMongoDB) Update(ctx context.Context, urlShortener domain.UrlShortener) error {
	filter := map[string]interface{}{
		"url_id": urlShortener.UrlId().Value(),
	}
	doc := NewUrlShortenerMongo(urlShortener)
	err := u.baseRepository.Update(ctx, filter, doc)
	if err != nil {
		return err
	}

	return nil
}
