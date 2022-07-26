package mother

import (
	vo2 "github.com/creepzed/url-shortener-service/app/shared/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
)

func UrlId(urlId string) vo2.UrlId {
	vo, _ := vo2.NewUrlId(urlId)
	return vo
}

func UrlEnabled(enabled bool) vo.UrlEnabled {
	return vo.NewUrlEnabled(enabled)
}

func OriginalUrl(originalUrl string) vo.OriginalUrl {
	vo, _ := vo.NewOriginalUrl(originalUrl)
	return vo
}

func UserId(userId string) vo.UserId {
	vo, _ := vo.NewUserId(userId)
	return vo
}
