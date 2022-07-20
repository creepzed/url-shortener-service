package transformer

import "github.com/creepzed/url-shortener-service/app/shortener/domain"

type UrlShortenerTransformer interface {
	Transform(shortener domain.UrlShortener) (interface{}, error)
}
