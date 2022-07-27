package request

import (
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/utils"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
)

//UpdateRequest
type UpdateRequest struct {
	OriginalUrl string `json:"original_url,omitempty"` // Original Url
	IsEnabled   *bool  `json:"is_enabled,omitempty"`   // Enabled
}

func (r UpdateRequest) String() string {
	return utils.EntityToJson(r)
}

func RandomUpdateRequest(urlId string, enabled *bool) UpdateRequest {
	return UpdateRequest{
		IsEnabled:   enabled,
		OriginalUrl: randomvalues.RandomOriginalUrl(),
	}
}

func FailRequestUpdateWithWrongOriginalUrl(urlId string, enabled *bool) UpdateRequest {
	return UpdateRequest{
		IsEnabled:   enabled,
		OriginalUrl: randomvalues.InvalidOriginalUrl(),
	}
}
func FailRequestUpdateWithWrongUserId(urlId string, enabled *bool) UpdateRequest {
	return UpdateRequest{
		IsEnabled:   enabled,
		OriginalUrl: randomvalues.RandomOriginalUrl(),
	}
}
