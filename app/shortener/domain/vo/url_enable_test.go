package vo

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUrlEnable(t *testing.T) {
	t.Parallel()
	t.Run("given a Enabled value, UrlEnable is true", func(t *testing.T) {
		value := Enabled

		vo := NewUrlEnabled(value)
		assert.Equal(t, vo.Value(), value)
	})

	t.Run("given a Disabled value, UrlEnable is true", func(t *testing.T) {
		value := Disabled

		vo := NewUrlEnabled(value)
		assert.Equal(t, vo.Value(), Disabled)
	})

}
