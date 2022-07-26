package finding

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/mocks/storagemocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/mocks/transformermocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGelAllApplicationService(t *testing.T) {
	t.Parallel()
	t.Run("given an invalid url, it should return the invalid url error", func(t *testing.T) {

		query := NewFindUrlShortenerQuery(randomvalues.InvalidUrlId(), nil)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)

		mockTransformer := transformermocks.NewUrlShortenerTransformer(t)

		service := NewFindApplicationService(mockRepository, mockTransformer)
		result, err := service.Do(context.Background(), query)

		assert.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrInvalidUrlId)
		assert.Nil(t, result)
	})
	t.Run("given a valid url, not found returns error", func(t *testing.T) {

		urlId := randomvalues.RandomUrlId()
		query := NewFindUrlShortenerQuery(urlId, nil)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(domain.UrlShortener{}, nil)

		mockTransformer := transformermocks.NewUrlShortenerTransformer(t)

		service := NewFindApplicationService(mockRepository, mockTransformer)
		result, err := service.Do(context.Background(), query)

		assert.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrUrlNotFound)
		assert.Equal(t, nil, result)
	})
	t.Run("given a valid url, return an error in the database", func(t *testing.T) {

		urlId := randomvalues.RandomUrlId()
		query := NewFindUrlShortenerQuery(urlId, nil)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(domain.UrlShortener{}, exception.ErrDataBase)

		mockTransformer := transformermocks.NewUrlShortenerTransformer(t)

		service := NewFindApplicationService(mockRepository, mockTransformer)
		result, err := service.Do(context.Background(), query)

		assert.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrDataBase)
		assert.Equal(t, nil, result)
	})
	t.Run("given a valid url, return an error in the transformer", func(t *testing.T) {
		urlId := randomvalues.RandomUrlId()
		query := NewFindUrlShortenerQuery(urlId, nil)
		urlExpected := domain.RandomUrlShortener(urlId, vo.Enabled)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(urlExpected, nil)

		mockTransformer := transformermocks.NewUrlShortenerTransformer(t)
		mockTransformer.
			On("Transform", mock.AnythingOfType("domain.UrlShortener")).
			Return(nil, exception.ErrTransforming)

		service := NewFindApplicationService(mockRepository, mockTransformer)
		result, err := service.Do(context.Background(), query)

		assert.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrTransforming)
		assert.Equal(t, nil, result)
	})
	t.Run("given a valid url, return data", func(t *testing.T) {
		urlId := randomvalues.RandomUrlId()
		query := NewFindUrlShortenerQuery(urlId, nil)
		urlExpected := domain.RandomUrlShortener(urlId, vo.Enabled)
		responseExpected := struct {
			UrlId       string `json:"url_id"`
			IsEnabled   bool   `json:"is_enabled"`
			OriginalUrl string `json:"original_url"`
		}{
			UrlId:       urlExpected.UrlId().Value(),
			IsEnabled:   urlExpected.IsEnabled().Value(),
			OriginalUrl: urlExpected.OriginalUrl().Value(),
		}

		mockTransformer := transformermocks.NewUrlShortenerTransformer(t)
		mockTransformer.
			On("Transform", mock.AnythingOfType("domain.UrlShortener")).
			Return(responseExpected, nil)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(urlExpected, nil)

		service := NewFindApplicationService(mockRepository, mockTransformer)
		result, err := service.Do(context.Background(), query)

		require.NoError(t, err)

		assert.Equal(t, responseExpected, result)
	})
}
