package getting

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/application/mocks/querymocks"
	"github.com/creepzed/url-shortener-service/app/shared/application/query"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/mocks/storagemocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/mocks/transformermocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetAllUrlShortenerQueryHandler(t *testing.T) {
	t.Parallel()
	t.Run("given a valid registered query it is executed", func(t *testing.T) {

		userId := randomvalues.RandomUserId()
		query := NewGetAllUrlShortenerQuery(userId)
		urlExpected := domain.RandomUrlShortener(randomvalues.RandomUrlId(), vo.Enabled)
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

		mockTransformer := transformermocks.NewUrlShortenerTransformer(t)
		mockTransformer.
			On("TransformList", mock.AnythingOfType("[]domain.UrlShortener")).
			Return(responseExpected, nil)

		repositoryMock := storagemocks.NewGetAllByUserIdRepository(t)

		repositoryMock.
			On("GetAllByUserId", context.Background(), mock.AnythingOfType("vo.UserId")).
			Return(listExpected, nil)

		service := NewGetAllApplicationService(repositoryMock, mockTransformer)

		handler := NewGetAllUrlShortenerQueryHandler(service)
		result, err := handler.Handle(context.Background(), query)

		assert.NoError(t, err)
		assert.Equal(t, responseExpected, result)
	})

	t.Run("given a valid unregistered query, return an error", func(t *testing.T) {

		var queryMockType query.Type = "query.mock"
		queryMock := querymocks.NewQuery(t)
		queryMock.On("Type").Return(queryMockType)

		repositoryMock := storagemocks.NewGetAllByUserIdRepository(t)
		mockTransformer := transformermocks.NewUrlShortenerTransformer(t)

		service := NewGetAllApplicationService(repositoryMock, mockTransformer)

		handler := NewGetAllUrlShortenerQueryHandler(service)
		result, err := handler.Handle(context.Background(), queryMock)

		require.Error(t, err)
		assert.ErrorIs(t, err, ErrUnexpectedQuery)
		assert.Nil(t, result)
	})
}
