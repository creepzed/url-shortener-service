package vo

import (
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUserId(t *testing.T) {
	t.Parallel()
	t.Run("give a valid value a userId is created", func(t *testing.T) {
		value := randomvalues.RandomUserId()

		vo, err := NewUserId(value)
		require.NoError(t, err)
		assert.Equal(t, vo.Value(), value)
	})

	t.Run("give a Invalid value a UserId is not created", func(t *testing.T) {
		value := randomvalues.InvalidUserId()

		_, err := NewUserId(value)
		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrInvalidUserId)
	})

	t.Run("given an empty value does not create a UserId", func(t *testing.T) {
		value := ""

		_, err := NewUserId(value)
		require.Error(t, err)
		assert.ErrorIs(t, err, exception.ErrInvalidUserId)
	})
}
