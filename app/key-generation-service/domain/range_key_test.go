package domain

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestNewRangeKey(t *testing.T) {
	t.Parallel()
	t.Run("should return a key", func(t *testing.T) {
		startRange := uint64(1)
		endRange := uint64(1)

		keys := NewRangeKey(startRange, endRange)
		key, err := keys.GetKey()

		require.NoError(t, err)
		require.NotEmpty(t, key)
	})

	t.Run("should return an error", func(t *testing.T) {
		startRange := uint64(1)
		endRange := uint64(1)

		keys := NewRangeKey(startRange, endRange)
		keys.GetKey()
		key, err := keys.GetKey()

		assert.Error(t, err)
		assert.Empty(t, key)
	})
}
