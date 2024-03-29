package creating

import (
	"context"
	"errors"
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/application/event"
	vo2 "github.com/creepzed/url-shortener-service/app/shared/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/repository"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
)

type CreateApplicationService interface {
	Do(ctx context.Context, command CreateUrlShortenerCommand) error
}

type createApplicationService struct {
	repository repository.CreateAndFindRepository
	eventBus   event.EventBus
}

func NewCreateApplicationService(repository repository.CreateAndFindRepository, eventBus event.EventBus) *createApplicationService {
	return &createApplicationService{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (cas createApplicationService) Do(ctx context.Context, command CreateUrlShortenerCommand) (err error) {

	urlId, err := vo2.NewUrlId(command.UrlId())
	if err != nil {
		return err
	}

	existUrl, err := cas.repository.FindById(ctx, urlId)
	if err != nil {
		if !errors.Is(err, exception.ErrUrlNotFound) {
			return fmt.Errorf("%w :%s", exception.ErrDataBase, err.Error())
		}
	}

	if existUrl.UrlId().Value() != "" {
		return fmt.Errorf("%w: %s", exception.ErrUrlIdDuplicate, urlId.Value())
	}

	originalUrl, err := vo.NewOriginalUrl(command.OriginalUrl())
	if err != nil {
		return err
	}

	userId, err := vo.NewUserId(command.UserId())
	if err != nil {
		return err
	}

	urlShortener := domain.NewUrlShortener(urlId, originalUrl, userId)

	err = cas.repository.Create(ctx, urlShortener)
	if err != nil {
		return fmt.Errorf("%w: %s", exception.ErrDataBase, err.Error())
	}

	err = cas.eventBus.Publish(ctx, urlShortener.PullEvents())
	if err != nil {
		return fmt.Errorf("%w: %s", exception.ErrEventBus, err.Error())
	}

	return
}
