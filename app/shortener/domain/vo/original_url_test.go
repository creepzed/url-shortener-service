package vo

import (
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestOriginalUrl(t *testing.T) {
	t.Parallel()
	t.Run("give a valid value a OriginalUrl is created", func(t *testing.T) {
		value := randomvalues.RandomOriginalUrl()

		vo, err := NewOriginalUrl(value)
		require.NoError(t, err)
		assert.Equal(t, vo.Value(), value)
	})

	t.Run("give a valid value a OriginalUrl is not created", func(t *testing.T) {
		value := randomvalues.InvalidOriginalUrl()

		_, err := NewOriginalUrl(value)
		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrInvalidOriginalUrl)
	})

	t.Run("given an empty value does not create a OriginalUrl", func(t *testing.T) {
		value := ""

		_, err := NewOriginalUrl(value)
		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrEmptyOriginalUrl)
	})
}
