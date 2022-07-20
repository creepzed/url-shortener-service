package request

import (
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/utils"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
)

type UrlShortenerCreateRequest struct {
	OriginalUrl string `json:"original_url" validate:"url,required"`
	UserId      string `json:"user_id" validate:"email,required"`
}

func (r UrlShortenerCreateRequest) String() string {
	return utils.EntityToJson(r)
}

func RandomUrlShortenerCreateRequest() UrlShortenerCreateRequest {
	return UrlShortenerCreateRequest{
		OriginalUrl: randomvalues.RandomOriginalUrl(),
		UserId:      randomvalues.RandomUserId(),
	}
}

func FailRequestWithWrongOriginalUrl() UrlShortenerCreateRequest {
	return UrlShortenerCreateRequest{
		OriginalUrl: randomvalues.InvalidOriginalUrl(),
		UserId:      randomvalues.RandomUserId(),
	}
}
func FailRequestWithWrongUserId() UrlShortenerCreateRequest {
	return UrlShortenerCreateRequest{
		OriginalUrl: randomvalues.RandomOriginalUrl(),
		UserId:      randomvalues.InvalidUserId(),
	}
}

func InvalidRequest() interface{} {
	return ""
}
