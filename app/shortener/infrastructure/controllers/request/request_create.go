package request

import (
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/utils"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
)

//UrlShortenerRequestCreate
type UrlShortenerRequestCreate struct {
	OriginalUrl string `json:"original_url" validate:"url,required"` // Original Url
	UserId      string `json:"user_id" validate:"email,required"`    // User Id
}

func (r UrlShortenerRequestCreate) String() string {
	return utils.EntityToJson(r)
}

func SetUrlShortenerRequestCreate(urlShortener domain.UrlShortener) UrlShortenerRequestCreate {
	return UrlShortenerRequestCreate{
		OriginalUrl: urlShortener.OriginalUrl().Value(),
		UserId:      urlShortener.OriginalUrl().Value(),
	}
}

func RandomUrlShortenerRequestCreate() UrlShortenerRequestCreate {
	return UrlShortenerRequestCreate{
		OriginalUrl: randomvalues.RandomOriginalUrl(),
		UserId:      randomvalues.RandomUserId(),
	}
}

func FailRequestCreateWithWrongOriginalUrl() UrlShortenerRequestCreate {
	return UrlShortenerRequestCreate{
		OriginalUrl: randomvalues.InvalidOriginalUrl(),
		UserId:      randomvalues.RandomUserId(),
	}
}
func FailRequestCreateWithWrongUserId() UrlShortenerRequestCreate {
	return UrlShortenerRequestCreate{
		OriginalUrl: randomvalues.RandomOriginalUrl(),
		UserId:      randomvalues.InvalidUserId(),
	}
}

func InvalidRequest() interface{} {
	return ""
}
