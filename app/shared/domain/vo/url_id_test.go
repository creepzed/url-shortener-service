package vo

import (
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUrlId(t *testing.T) {
	t.Parallel()
	t.Run("give a valid value a UrlId is created", func(t *testing.T) {
		value := randomvalues.RandomUrlId()

		vo, err := NewUrlId(value)
		require.NoError(t, err)
		assert.Equal(t, vo.Value(), value)
	})

	t.Run("give a valid value a UrlId is not created", func(t *testing.T) {
		value := randomvalues.InvalidUrlId()

		_, err := NewUrlId(value)
		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrInvalidUrlId)
	})

	t.Run("given an empty value does not create a UrlId", func(t *testing.T) {
		value := ""

		_, err := NewUrlId(value)
		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrEmptyUrlId)
	})
}
