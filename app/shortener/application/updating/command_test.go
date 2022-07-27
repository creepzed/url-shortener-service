package updating

import (
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUpdateUrlShortenerCommand(t *testing.T) {
	t.Parallel()
	t.Run("given a valid request, the command is built", func(t *testing.T) {
		urlId := randomvalues.RandomUrlId()
		isEnabled := randomvalues.RandomIsEnabled()
		originalUrl := randomvalues.RandomOriginalUrl()

		commandType := UpdateUrlShortenerCommandType

		cmd := NewUpdateUrlShortenerCommand(urlId, &isEnabled, originalUrl)

		assert.Equal(t, cmd.Type(), commandType)
		assert.Equal(t, cmd.UrlId(), urlId)
		auxIsEnabled := cmd.IsEnabled()
		assert.Equal(t, *auxIsEnabled, isEnabled)
		assert.Equal(t, cmd.OriginalUrl(), originalUrl)
	})
}
