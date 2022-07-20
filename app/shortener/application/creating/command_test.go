package creating

import (
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateUrlShortenerCommand(t *testing.T) {
	t.Parallel()
	t.Run("given a valid request, the command is built", func(t *testing.T) {
		urlId := randomvalues.RandomUrlId()
		originalUrl := randomvalues.RandomOriginalUrl()
		userId := randomvalues.RandomUserId()
		commandType := CreateUrlShortenerCommandType

		cmd := NewCreateUrlShortenerCommand(urlId, originalUrl, userId)

		assert.Equal(t, cmd.Type(), commandType)
		assert.Equal(t, cmd.UrlId(), urlId)
		assert.Equal(t, cmd.OriginalUrl(), originalUrl)
		assert.Equal(t, cmd.UserId(), userId)
	})
}
