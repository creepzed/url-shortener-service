package creating

import (
	"github.com/creepzed/url-shortener-service/app/shared/application/command"
)

const CreateUrlShortenerCommandType command.Type = "shortener.create.url"

type CreateUrlShortenerCommand struct {
	urlId       string
	originalUrl string
	userId      string
}

func NewCreateUrlShortenerCommand(urlId string, originalUrl string, userId string) CreateUrlShortenerCommand {
	return CreateUrlShortenerCommand{
		urlId:       urlId,
		originalUrl: originalUrl,
		userId:      userId,
	}
}

func (c CreateUrlShortenerCommand) UrlId() string {
	return c.urlId
}

func (c CreateUrlShortenerCommand) OriginalUrl() string {
	return c.originalUrl
}

func (c CreateUrlShortenerCommand) UserId() string {
	return c.userId
}

func (c CreateUrlShortenerCommand) Type() command.Type {
	return CreateUrlShortenerCommandType
}
