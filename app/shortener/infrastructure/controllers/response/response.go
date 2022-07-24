package response

type OutputResponse struct {
	UrlId       string `json:"url_id"`
	IsEnabled   bool   `json:"is_enabled"`
	OriginalUrl string `json:"original_url"`
	UserId      string `json:"user_id"`
}
