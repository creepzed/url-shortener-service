package mother

import "github.com/creepzed/url-shortener-service/app/shortener/domain/vo"

func UrlId(urlId string) vo.UrlId {
	vo, _ := vo.NewUrlId(urlId)
	return vo
}

func OriginalUrl(originalUrl string) vo.OriginalUrl {
	vo, _ := vo.NewOriginalUrl(originalUrl)
	return vo
}

func UserId(userId string) vo.UserId {
	vo, _ := vo.NewUserId(userId)
	return vo
}
