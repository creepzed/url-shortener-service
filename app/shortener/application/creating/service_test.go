package creating

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/application/mocks/eventmocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/mocks/storagemocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/mother"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateApplicationService(t *testing.T) {
	t.Parallel()

	t.Run("Receiving a create URL shortener command, but the UrlId is not valid", func(t *testing.T) {
		command := NewCreateUrlShortenerCommand(randomvalues.InvalidUrlId(), randomvalues.RandomOriginalUrl(), randomvalues.RandomUserId())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)

		mockEventBus := eventmocks.NewEventBus(t)

		service := NewCreateApplicationService(mockRepository, mockEventBus)

		err := service.Do(context.Background(), command)

		assert.ErrorIs(t, err, exception.ErrInvalidUrlId)
	})

	t.Run("Receiving a create URL shortener command, but the originalUrl is not valid", func(t *testing.T) {
		command := NewCreateUrlShortenerCommand(randomvalues.RandomUrlId(), randomvalues.InvalidOriginalUrl(), randomvalues.RandomUserId())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(domain.UrlShortener{}, nil)

		mockEventBus := eventmocks.NewEventBus(t)

		service := NewCreateApplicationService(mockRepository, mockEventBus)

		err := service.Do(context.Background(), command)

		assert.ErrorIs(t, err, exception.ErrInvalidOriginalUrl)
	})

	t.Run("Receiving a create URL shortener command, but the UserId is not valid", func(t *testing.T) {
		command := NewCreateUrlShortenerCommand(randomvalues.RandomUrlId(), randomvalues.RandomOriginalUrl(), randomvalues.InvalidUserId())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(domain.UrlShortener{}, nil)

		mockEventBus := eventmocks.NewEventBus(t)

		service := NewCreateApplicationService(mockRepository, mockEventBus)

		err := service.Do(context.Background(), command)

		assert.ErrorIs(t, err, exception.ErrInvalidUserId)
	})

	t.Run("When receiving a valid create url shortener command is created, it returns nil", func(t *testing.T) {
		command := NewCreateUrlShortenerCommand(randomvalues.RandomUrlId(), randomvalues.RandomOriginalUrl(), randomvalues.RandomUserId())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("Create", context.Background(), mock.AnythingOfType("domain.UrlShortener")).
			Return(nil)

		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(domain.UrlShortener{}, nil)

		mockEventBus := eventmocks.NewEventBus(t)
		mockEventBus.
			On("Publish", context.Background(), mock.AnythingOfType("[]event.Event")).
			Return(nil)

		service := NewCreateApplicationService(mockRepository, mockEventBus)

		err := service.Do(context.Background(), command)

		assert.Nil(t, err)
	})

	t.Run("When receiving a valid create url shortener, but UrlId is duplicate", func(t *testing.T) {

		urlId := randomvalues.RandomUrlId()
		originalUrl := randomvalues.RandomOriginalUrl()
		userId := randomvalues.RandomUserId()

		existUrl := domain.NewUrlShortener(mother.UrlId(urlId), mother.OriginalUrl(originalUrl), mother.UserId(userId))

		command := NewCreateUrlShortenerCommand(urlId, originalUrl, userId)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)

		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(existUrl, nil)

		mockEventBus := eventmocks.NewEventBus(t)

		service := NewCreateApplicationService(mockRepository, mockEventBus)

		err := service.Do(context.Background(), command)

		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrUrlIdDuplicate)
	})

	t.Run("When receiving a valid create url shortener command, but return an create error on database ", func(t *testing.T) {
		command := NewCreateUrlShortenerCommand(randomvalues.RandomUrlId(), randomvalues.RandomOriginalUrl(), randomvalues.RandomUserId())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("Create", context.Background(), mock.AnythingOfType("domain.UrlShortener")).
			Return(exception.ErrDataBase)

		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(domain.UrlShortener{}, nil)

		mockEventBus := eventmocks.NewEventBus(t)

		service := NewCreateApplicationService(mockRepository, mockEventBus)

		err := service.Do(context.Background(), command)

		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrDataBase)
	})

	t.Run("When receiving a valid create url shortener command, but return an find error on database ", func(t *testing.T) {
		command := NewCreateUrlShortenerCommand(randomvalues.RandomUrlId(), randomvalues.RandomOriginalUrl(), randomvalues.RandomUserId())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
	
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(domain.UrlShortener{}, exception.ErrDataBase)

		mockEventBus := eventmocks.NewEventBus(t)

		service := NewCreateApplicationService(mockRepository, mockEventBus)

		err := service.Do(context.Background(), command)

		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrDataBase)
	})

	t.Run("When receiving a valid create url shortener command, but return an save error on eventbus ", func(t *testing.T) {
		command := NewCreateUrlShortenerCommand(randomvalues.RandomUrlId(), randomvalues.RandomOriginalUrl(), randomvalues.RandomUserId())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("Create", context.Background(), mock.AnythingOfType("domain.UrlShortener")).
			Return(nil)

		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(domain.UrlShortener{}, nil)

		mockEventBus := eventmocks.NewEventBus(t)
		mockEventBus.
			On("Publish", context.Background(), mock.AnythingOfType("[]event.Event")).
			Return(exception.ErrEventBus)

		service := NewCreateApplicationService(mockRepository, mockEventBus)

		err := service.Do(context.Background(), command)

		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrEventBus)
	})
}
