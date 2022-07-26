package generating

import (
	"errors"
	"github.com/creepzed/url-shortener-service/app/key-generation-service/domain"
	"github.com/creepzed/url-shortener-service/app/key-generation-service/domain/mocks/rangekeymocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestKeyGenerateService(t *testing.T) {
	t.Parallel()
	t.Run("when the repository returns an error", func(t *testing.T) {
		rangeRepositoryMock := rangekeymocks.NewRangeKeyRepository(t)
		rangeRepositoryMock.On("GetRange").
			Return(nil, errors.New("an error"))

		service, err := NewKeyGenerateService(rangeRepositoryMock)
		assert.Error(t, err)
		assert.Nil(t, service)
	})

	t.Run("when the repository returns an error", func(t *testing.T) {
		rangeKey := domain.NewRangeKey(1, 100)

		rangeRepositoryMock := rangekeymocks.NewRangeKeyRepository(t)
		rangeRepositoryMock.On("GetRange").
			Return(rangeKey, nil)

		service, err := NewKeyGenerateService(rangeRepositoryMock)
		require.NoError(t, err)

		key, err := service.GetKey()
		assert.NoError(t, err)
		assert.NotEmpty(t, key)
	})
}
