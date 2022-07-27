package response

//OutputResponse
type OutputResponse struct {
	UrlId       string `json:"url_id"`       //Url Id
	IsEnabled   bool   `json:"is_enabled"`   //Is Enabled
	OriginalUrl string `json:"original_url"` //Original Url
	UserId      string `json:"user_id"`      //User Id
}
