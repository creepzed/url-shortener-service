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

func (t transformer) TransformList(shortenerList []domain.UrlShortener) ([]interface{}, error) {
	result := make([]interface{}, 0)
	for _, urlShortener := range shortenerList {
		short, err := t.Transform(urlShortener)
		if err != nil {
			return nil, err
		}
		result = append(result, short)
	}
	return result, nil
}
