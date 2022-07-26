package getting

import "github.com/creepzed/url-shortener-service/app/shared/application/query"

const GetAllUrlShortenerQueryType query.Type = "shortener.get.all.url"

type GetAllUrlShortenerQuery struct {
	userId string
}

func NewGetAllUrlShortenerQuery(userId string) GetAllUrlShortenerQuery {
	return GetAllUrlShortenerQuery{
		userId: userId,
	}
}

func (f GetAllUrlShortenerQuery) UserId() string {
	return f.userId
}

func (f GetAllUrlShortenerQuery) Type() query.Type {
	return GetAllUrlShortenerQueryType
}
