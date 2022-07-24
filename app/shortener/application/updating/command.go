package updating

import (
	"github.com/creepzed/url-shortener-service/app/shared/application/command"
)

const UpdateUrlShortenerCommandType command.Type = "shortener.update.url"

type UpdateUrlShortenerCommand struct {
	urlId       string
	isEnabled   *bool
	originalUrl string
	userId      string
}

func NewUpdateUrlShortenerCommand(urlId string, isEnabled *bool, originalUrl string, userId string) UpdateUrlShortenerCommand {
	return UpdateUrlShortenerCommand{
		urlId:       urlId,
		isEnabled:   isEnabled,
		originalUrl: originalUrl,
		userId:      userId,
	}
}

func (c UpdateUrlShortenerCommand) UrlId() string {
	return c.urlId
}

func (c UpdateUrlShortenerCommand) IsEnabled() *bool {
	return c.isEnabled
}

func (c UpdateUrlShortenerCommand) OriginalUrl() string {
	return c.originalUrl
}

func (c UpdateUrlShortenerCommand) UserId() string {
	return c.userId
}

func (c UpdateUrlShortenerCommand) Type() command.Type {
	return UpdateUrlShortenerCommandType
}
