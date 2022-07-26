package updating

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/application/mocks/eventmocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/mocks/storagemocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUpdateApplicationService(t *testing.T) {
	t.Parallel()
	t.Run("given an invalid url, it should return the invalid url error", func(t *testing.T) {

		auxIsEnabled := randomvalues.RandomIsEnabled()
		cmd := NewUpdateUrlShortenerCommand(randomvalues.InvalidUrlId(), &auxIsEnabled,
			randomvalues.RandomOriginalUrl(), randomvalues.RandomUserId())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)

		mockEventBus := eventmocks.NewEventBus(t)

		service := NewUpdateApplicationService(mockRepository, mockEventBus)
		err := service.Do(context.Background(), cmd)

		assert.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrInvalidUrlId)

	})
	t.Run("given a valid url, not found returns error", func(t *testing.T) {

		auxIsEnabled := randomvalues.RandomIsEnabled()
		cmd := NewUpdateUrlShortenerCommand(randomvalues.RandomUrlId(), &auxIsEnabled,
			randomvalues.RandomOriginalUrl(), randomvalues.RandomUserId())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(domain.UrlShortener{}, nil)

		mockEventBus := eventmocks.NewEventBus(t)

		service := NewUpdateApplicationService(mockRepository, mockEventBus)
		err := service.Do(context.Background(), cmd)

		assert.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrUrlNotFound)
	})
	t.Run("given a valid url, it returns an error in the database when it is finding", func(t *testing.T) {

		auxIsEnabled := randomvalues.RandomIsEnabled()
		cmd := NewUpdateUrlShortenerCommand(randomvalues.RandomUrlId(), &auxIsEnabled,
			randomvalues.RandomOriginalUrl(), randomvalues.RandomUserId())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(domain.UrlShortener{}, exception.ErrDataBase)

		mockEventBus := eventmocks.NewEventBus(t)

		service := NewUpdateApplicationService(mockRepository, mockEventBus)
		err := service.Do(context.Background(), cmd)

		assert.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrDataBase)
	})
	t.Run("given a valid url, it returns an error in the database when it is updating", func(t *testing.T) {

		urlId := randomvalues.RandomUrlId()
		urlExpected := domain.RandomUrlShortener(urlId, vo.Enabled)
		auxIsEnabled := vo.Disabled
		cmd := NewUpdateUrlShortenerCommand(urlExpected.UrlId().Value(),
			&auxIsEnabled, urlExpected.OriginalUrl().Value(), urlExpected.UserId().Value())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(urlExpected, nil)
		mockRepository.
			On("Update", context.Background(), mock.AnythingOfType("domain.UrlShortener")).
			Return(exception.ErrDataBase)

		mockEventBus := eventmocks.NewEventBus(t)

		service := NewUpdateApplicationService(mockRepository, mockEventBus)
		err := service.Do(context.Background(), cmd)

		assert.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrDataBase)
	})

	t.Run("given a valid url, it returns an error in the event bus when it is updating", func(t *testing.T) {
		urlId := randomvalues.RandomUrlId()
		urlExpected := domain.RandomUrlShortener(urlId, vo.Enabled)
		auxIsEnabled := vo.Disabled
		cmd := NewUpdateUrlShortenerCommand(urlExpected.UrlId().Value(),
			&auxIsEnabled, urlExpected.OriginalUrl().Value(), urlExpected.UserId().Value())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(urlExpected, nil)
		mockRepository.
			On("Update", context.Background(), mock.AnythingOfType("domain.UrlShortener")).
			Return(nil)

		mockEventBus := eventmocks.NewEventBus(t)
		mockEventBus.
			On("Publish", context.Background(), mock.AnythingOfType("[]event.Event")).
			Return(exception.ErrEventBus)

		service := NewUpdateApplicationService(mockRepository, mockEventBus)
		err := service.Do(context.Background(), cmd)

		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrEventBus)
	})

	t.Run("given a valid url, when invalid the original url returns an error", func(t *testing.T) {
		urlId := randomvalues.RandomUrlId()
		urlExpected := domain.RandomUrlShortener(urlId, vo.Enabled)
		auxIsEnabled := vo.Disabled
		cmd := NewUpdateUrlShortenerCommand(urlExpected.UrlId().Value(),
			&auxIsEnabled, randomvalues.InvalidOriginalUrl(), urlExpected.UserId().Value())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(urlExpected, nil)

		mockEventBus := eventmocks.NewEventBus(t)

		service := NewUpdateApplicationService(mockRepository, mockEventBus)
		err := service.Do(context.Background(), cmd)

		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrInvalidOriginalUrl)
	})

	t.Run("given a valid url, when invalid the userid returns an error", func(t *testing.T) {
		urlId := randomvalues.RandomUrlId()
		urlExpected := domain.RandomUrlShortener(urlId, vo.Enabled)
		auxIsEnabled := vo.Disabled
		cmd := NewUpdateUrlShortenerCommand(urlExpected.UrlId().Value(),
			&auxIsEnabled, urlExpected.OriginalUrl().Value(), randomvalues.InvalidUrlId())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(urlExpected, nil)

		mockEventBus := eventmocks.NewEventBus(t)

		service := NewUpdateApplicationService(mockRepository, mockEventBus)
		err := service.Do(context.Background(), cmd)

		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrInvalidUserId)
	})

	t.Run("given a valid url, return ok", func(t *testing.T) {
		urlId := randomvalues.RandomUrlId()
		urlExpected := domain.RandomUrlShortener(urlId, vo.Enabled)
		auxIsEnabled := vo.Disabled
		cmd := NewUpdateUrlShortenerCommand(urlExpected.UrlId().Value(),
			&auxIsEnabled, urlExpected.OriginalUrl().Value(), urlExpected.UserId().Value())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(urlExpected, nil)
		mockRepository.
			On("Update", context.Background(), mock.AnythingOfType("domain.UrlShortener")).
			Return(nil)

		mockEventBus := eventmocks.NewEventBus(t)
		mockEventBus.
			On("Publish", context.Background(), mock.AnythingOfType("[]event.Event")).
			Return(nil)

		service := NewUpdateApplicationService(mockRepository, mockEventBus)
		err := service.Do(context.Background(), cmd)

		require.NoError(t, err)
	})
	//TODO testing to optimistic locking
}
