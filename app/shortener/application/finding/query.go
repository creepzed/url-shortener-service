package finding

import "github.com/creepzed/url-shortener-service/app/shared/application/query"

const FindUrlShortenerQueryType query.Type = "shortener.find.url"

type Metadata map[string]interface{}
type FindUrlShortenerQuery struct {
	urlId    string
	metadata Metadata
}

func NewFindUrlShortenerQuery(urlId string, metadata Metadata) FindUrlShortenerQuery {
	return FindUrlShortenerQuery{
		urlId:    urlId,
		metadata: metadata,
	}
}

func (f FindUrlShortenerQuery) UrlId() string {
	return f.urlId
}

func (f FindUrlShortenerQuery) Metadata() map[string]interface{} {
	return f.metadata
}

func (f FindUrlShortenerQuery) Type() query.Type {
	return FindUrlShortenerQueryType
}
