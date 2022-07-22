package domain

import (
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/mother"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
)

func RandomUrlShortener(urlId string, enabled bool) UrlShortener {
	return LoadUrlShortener(
		mother.UrlId(urlId),
		mother.UrlEnabled(enabled),
		mother.OriginalUrl(randomvalues.RandomOriginalUrl()),
		mother.UserId(randomvalues.RandomUserId()),
	)
}
