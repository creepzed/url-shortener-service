package transformer

import (
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/controllers/response"
)

type transformer struct {
}

func NewTransformer() transformer {
	return transformer{}
}

func (t transformer) Transform(urlShortener domain.UrlShortener) (interface{}, error) {

	anUrlShortener := response.OutputResponse{
		UrlId:       urlShortener.UrlId().Value(),
		IsEnabled:   urlShortener.IsEnabled().Value(),
		OriginalUrl: urlShortener.OriginalUrl().Value(),
		UserId:      urlShortener.UserId().Value(),
	}

	return anUrlShortener, nil
}
