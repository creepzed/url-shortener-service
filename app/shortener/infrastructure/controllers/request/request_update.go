package request

import (
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/utils"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
)

type UpdateRequest struct {
	OriginalUrl string `json:"original_url,omitempty"`
	IsEnabled   *bool  `json:"is_enabled,omitempty"`
	UserId      string `json:"user_id,omitempty"`
}

func (r UpdateRequest) String() string {
	return utils.EntityToJson(r)
}

func RandomUpdateRequest(urlId string, enabled *bool) UpdateRequest {
	return UpdateRequest{
		IsEnabled:   enabled,
		OriginalUrl: randomvalues.RandomOriginalUrl(),
		UserId:      randomvalues.RandomUserId(),
	}
}

func FailRequestUpdateWithWrongOriginalUrl(urlId string, enabled *bool) UpdateRequest {
	return UpdateRequest{
		IsEnabled:   enabled,
		OriginalUrl: randomvalues.InvalidOriginalUrl(),
		UserId:      randomvalues.RandomUserId(),
	}
}
func FailRequestUpdateWithWrongUserId(urlId string, enabled *bool) UpdateRequest {
	return UpdateRequest{
		IsEnabled:   enabled,
		OriginalUrl: randomvalues.RandomOriginalUrl(),
		UserId:      randomvalues.InvalidUserId(),
	}
}
