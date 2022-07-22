package finding

import "github.com/creepzed/url-shortener-service/app/shared/application/query"

const FindUrlShortenerQueryType query.Type = "shortener.find.url"

type FindUrlShortenerQuery struct {
	urlId string
}

func NewFindUrlShortenerQuery(urlId string) FindUrlShortenerQuery {
	return FindUrlShortenerQuery{
		urlId: urlId,
	}
}

func (f FindUrlShortenerQuery) UrlId() string {
	return f.urlId
}

func (f FindUrlShortenerQuery) Type() query.Type {
	return FindUrlShortenerQueryType
}
