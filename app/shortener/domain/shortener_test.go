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

	t.Run("given valid values UrlShortener is created and posted create event", func(t *testing.T) {
		urlId := mother.UrlId(randomvalues.RandomUrlId())
		originalUrl := mother.OriginalUrl(randomvalues.RandomOriginalUrl())
		userId := mother.UserId(randomvalues.RandomUserId())

		urlShortener := NewUrlShortener(urlId, originalUrl, userId)
		evt := urlShortener.PullEvents()[0]

		require.Equal(t, urlId.Value(), urlShortener.UrlId().Value())
		require.Equal(t, vo.Enabled, urlShortener.IsEnabled().Value())
		require.Equal(t, originalUrl.Value(), urlShortener.OriginalUrl().Value())
		require.Equal(t, userId.Value(), urlShortener.UserId().Value())
		assert.Equal(t, 1, len(urlShortener.PullEvents()))
		assert.Equal(t, true, urlShortener.IsChanged())
		assert.Equal(t, event.ShortenerCreatedEventType, evt.Type())
	})
	t.Run("given valid values UrlShortener is loaded, but does not post events", func(t *testing.T) {
		urlId := mother.UrlId(randomvalues.RandomUrlId())
		urlEnabled := vo.NewUrlEnabled(vo.Enabled)
		originalUrl := mother.OriginalUrl(randomvalues.RandomOriginalUrl())
		userId := mother.UserId(randomvalues.RandomUserId())

		urlShortener := LoadUrlShortener(urlId, urlEnabled, originalUrl, userId)

		require.Equal(t, urlId.Value(), urlShortener.UrlId().Value())
		require.Equal(t, urlEnabled.Value(), urlShortener.IsEnabled().Value())
		require.Equal(t, originalUrl.Value(), urlShortener.OriginalUrl().Value())
		require.Equal(t, userId.Value(), urlShortener.UserId().Value())
		assert.Equal(t, 0, len(urlShortener.PullEvents()))
		assert.Equal(t, false, urlShortener.IsChanged())
	})
	t.Run("given a loaded UrlShortener, when changed, post the update event", func(t *testing.T) {
		urlId := mother.UrlId(randomvalues.RandomUrlId())
		urlShortener := LoadUrlShortener(urlId, vo.NewUrlEnabled(vo.Enabled), mother.OriginalUrl(randomvalues.RandomOriginalUrl()), mother.UserId(randomvalues.RandomUserId()))

		urlEnabled := vo.NewUrlEnabled(!urlShortener.IsEnabled().Value())
		originalUrl := mother.OriginalUrl(randomvalues.RandomOriginalUrl())
		userId := mother.UserId(randomvalues.RandomUserId())

		urlShortener.Update(urlEnabled, originalUrl, userId)

		require.Equal(t, urlId.Value(), urlShortener.UrlId().Value())
		require.Equal(t, urlEnabled.Value(), urlShortener.IsEnabled().Value())
		require.Equal(t, originalUrl.Value(), urlShortener.OriginalUrl().Value())
		require.Equal(t, userId.Value(), urlShortener.UserId().Value())
		assert.Equal(t, 1, len(urlShortener.PullEvents()))
		assert.Equal(t, true, urlShortener.IsChanged())
	})
}
