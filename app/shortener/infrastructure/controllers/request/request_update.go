package request

import (
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/utils"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
)

type UpdateRequest struct {
	OriginalUrl string `json:"original_url,omitempty"`
	IsEnabled   bool   `json:"is_enabled,omitempty"`
	UserId      string `json:"user_id,omitempty"`
}

func (r UpdateRequest) String() string {
	return utils.EntityToJson(r)
}

func RandomUpdateRequest(urlId string, isEnabled bool) UpdateRequest {
	return UpdateRequest{
		IsEnabled:   isEnabled,
		OriginalUrl: randomvalues.RandomOriginalUrl(),
		UserId:      randomvalues.RandomUserId(),
	}
}

func FailRequestUpdateWithWrongOriginalUrl(urlId string, isEnabled bool) UpdateRequest {
	return UpdateRequest{
		IsEnabled:   isEnabled,
		OriginalUrl: randomvalues.InvalidOriginalUrl(),
		UserId:      randomvalues.RandomUserId(),
	}
}
func FailRequestUpdateWithWrongUserId(urlId string, isEnabled bool) UpdateRequest {
	return UpdateRequest{
		IsEnabled:   isEnabled,
		OriginalUrl: randomvalues.RandomOriginalUrl(),
		UserId:      randomvalues.InvalidUserId(),
	}
}
