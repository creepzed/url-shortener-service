package getting

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

func TestGetAllApplicationService(t *testing.T) {
	t.Parallel()
	t.Run("given an invalid userid, it should return the invalid user error", func(t *testing.T) {

		query := NewGetAllUrlShortenerQuery(randomvalues.InvalidUserId())

		mockRepository := storagemocks.NewGetAllByUserIdRepository(t)

		mockTransformer := transformermocks.NewUrlShortenerTransformer(t)

		service := NewGetAllApplicationService(mockRepository, mockTransformer)
		result, err := service.Do(context.Background(), query)

		assert.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrInvalidUserId)
		assert.Nil(t, result)
	})

	t.Run("given a valid userid, return an empty array", func(t *testing.T) {

		userId := randomvalues.RandomUserId()
		query := NewGetAllUrlShortenerQuery(userId)

		mockRepository := storagemocks.NewGetAllByUserIdRepository(t)
		mockRepository.
			On("GetAllByUserId", context.Background(), mock.AnythingOfType("vo.UserId")).
			Return([]domain.UrlShortener{}, nil)

		mockTransformer := transformermocks.NewUrlShortenerTransformer(t)
		mockTransformer.
			On("TransformList", mock.AnythingOfType("[]domain.UrlShortener")).
			Return([]interface{}{}, nil)

		service := NewGetAllApplicationService(mockRepository, mockTransformer)
		result, err := service.Do(context.Background(), query)

		assert.NoError(t, err)
		assert.Equal(t, []interface{}{}, result)
	})

	t.Run("given a valid userid, return an error in the database", func(t *testing.T) {

		userId := randomvalues.RandomUserId()
		query := NewGetAllUrlShortenerQuery(userId)

		mockRepository := storagemocks.NewGetAllByUserIdRepository(t)
		mockRepository.
			On("GetAllByUserId", context.Background(), mock.AnythingOfType("vo.UserId")).
			Return([]domain.UrlShortener{}, exception.ErrDataBase)

		mockTransformer := transformermocks.NewUrlShortenerTransformer(t)

		service := NewGetAllApplicationService(mockRepository, mockTransformer)
		result, err := service.Do(context.Background(), query)

		assert.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrDataBase)
		assert.Equal(t, nil, result)
	})
	t.Run("given a valid userid, return an error in the transformer", func(t *testing.T) {
		userId := randomvalues.RandomUserId()
		query := NewGetAllUrlShortenerQuery(userId)
		urlExpected := domain.RandomUrlShortener(userId, vo.Enabled)
		listExpected := []domain.UrlShortener{
			urlExpected,
		}

		mockRepository := storagemocks.NewGetAllByUserIdRepository(t)
		mockRepository.
			On("GetAllByUserId", context.Background(), mock.AnythingOfType("vo.UserId")).
			Return(listExpected, nil)

		mockTransformer := transformermocks.NewUrlShortenerTransformer(t)
		mockTransformer.
			On("TransformList", mock.AnythingOfType("[]domain.UrlShortener")).
			Return(nil, exception.ErrTransforming)

		service := NewGetAllApplicationService(mockRepository, mockTransformer)
		result, err := service.Do(context.Background(), query)

		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrTransforming)
		assert.Equal(t, nil, result)
	})
	t.Run("given a valid userid, return data", func(t *testing.T) {
		userId := randomvalues.RandomUserId()
		query := NewGetAllUrlShortenerQuery(userId)
		urlExpected := domain.RandomUrlShortener(userId, vo.Enabled)
		listExpected := []domain.UrlShortener{
			urlExpected,
		}
		responseExpected := []interface{}{
			struct {
				UrlId       string `json:"url_id"`
				IsEnabled   bool   `json:"is_enabled"`
				OriginalUrl string `json:"original_url"`
			}{
				UrlId:       urlExpected.UrlId().Value(),
				IsEnabled:   urlExpected.IsEnabled().Value(),
				OriginalUrl: urlExpected.OriginalUrl().Value(),
			},
		}

		mockRepository := storagemocks.NewGetAllByUserIdRepository(t)
		mockRepository.
			On("GetAllByUserId", context.Background(), mock.AnythingOfType("vo.UserId")).
			Return(listExpected, nil)

		mockTransformer := transformermocks.NewUrlShortenerTransformer(t)
		mockTransformer.
			On("TransformList", mock.AnythingOfType("[]domain.UrlShortener")).
			Return(responseExpected, nil)

		service := NewGetAllApplicationService(mockRepository, mockTransformer)
		result, err := service.Do(context.Background(), query)

		require.NoError(t, err)

		assert.Equal(t, responseExpected, result)
	})
}
