package updating

import (
	"context"
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/application/event"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/repository"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
)

type UpdateApplicationService interface {
	Do(ctx context.Context, command UpdateUrlShortenerCommand) error
}

type updateApplicationService struct {
	repository repository.UpdateAndFindRepository
	eventBus   event.EventBus
}

func NewUpdateApplicationService(repository repository.UpdateAndFindRepository, eventBus event.EventBus) *updateApplicationService {
	return &updateApplicationService{
		repository: repository,
		eventBus:   eventBus,
	}
}

func (uas updateApplicationService) Do(ctx context.Context, command UpdateUrlShortenerCommand) (err error) {
	urlId, err := vo.NewUrlId(command.UrlId())
	if err != nil {
		return err
	}

	urlShortener, err := uas.repository.FindById(ctx, urlId)
	if err != nil {
		return fmt.Errorf("%w: %s", exception.ErrDataBase, err.Error())
	}

	if urlShortener.UrlId().Value() != urlId.Value() {
		return fmt.Errorf("%w: %s", exception.ErrUrlNotFound, urlId.Value())
	}

	isEnabled := urlShortener.IsEnabled()
	if command.IsEnabled() != nil {
		value := *command.IsEnabled()
		isEnabled = vo.NewUrlEnabled(value)
	}

	originalUrl := urlShortener.OriginalUrl()
	if len(command.OriginalUrl()) > 0 {
		originalUrl, err = vo.NewOriginalUrl(command.OriginalUrl())
		if err != nil {
			return err
		}
	}

	userId := urlShortener.UserId()
	if len(command.UserId()) > 0 {
		userId, err = vo.NewUserId(command.UserId())
		if err != nil {
			return err
		}
	}

	urlShortener.Update(isEnabled, originalUrl, userId)

	err = uas.repository.Update(ctx, urlShortener)
	if err != nil {
		return err
	}

	err = uas.eventBus.Publish(ctx, urlShortener.PullEvents())
	if err != nil {
		return fmt.Errorf("%w: %s", exception.ErrEventBus, err.Error())
	}

	return
}
