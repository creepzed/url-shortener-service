package domain

import (
	"github.com/creepzed/url-shortener-service/app/shortener/domain/event"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/mother"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUrlShortener(t *testing.T) {
	t.Parallel()

	t.Run("given valid values UrlShortener is created and posted event created", func(t *testing.T) {
		urlId := mother.UrlId(randomvalues.RandomUrlId())
		originalUrl := mother.OriginalUrl(randomvalues.RandomOriginalUrl())
		userId := mother.UserId(randomvalues.RandomUserId())

		urlShortener := NewUrlShortener(urlId, originalUrl, userId)
		evt := urlShortener.PullEvents()[0]

		require.Equal(t, urlId.Value(), urlShortener.UrlId().Value())
		require.Equal(t, originalUrl.Value(), urlShortener.OriginalUrl().Value())
		require.Equal(t, userId.Value(), urlShortener.UserId().Value())
		assert.Equal(t, 1, len(urlShortener.PullEvents()))
		assert.Equal(t, event.ShortenerCreatedEventType, evt.Type())
	})

	t.Run("when UrlShortener loads, but does not raise events", func(t *testing.T) {
		urlId := mother.UrlId(randomvalues.RandomUrlId())
		urlEnabled := vo.NewUrlEnabled(vo.Enabled)
		originalUrl := mother.OriginalUrl(randomvalues.RandomOriginalUrl())
		userId := mother.UserId(randomvalues.RandomUserId())

		urlShortener := LoadUrlShortener(urlId, urlEnabled, originalUrl, userId)

		require.Equal(t, urlId.Value(), urlShortener.UrlId().Value())
		require.Equal(t, originalUrl.Value(), urlShortener.OriginalUrl().Value())
		require.Equal(t, userId.Value(), urlShortener.UserId().Value())
		assert.Equal(t, 0, len(urlShortener.PullEvents()))
	})
}
