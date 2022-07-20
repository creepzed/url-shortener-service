package creating

import (
	"context"
	"errors"
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/application/command"
)

type CreateUrlShortenerCommandHandler struct {
	applicationService CreateApplicationService
}

var (
	ErrUnexpectedCommand = errors.New("unexpected command")
)

func NewCreateUrlShortenerCommandHandler(applicationService CreateApplicationService) CreateUrlShortenerCommandHandler {
	return CreateUrlShortenerCommandHandler{
		applicationService: applicationService,
	}
}

func (h CreateUrlShortenerCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	command, ok := cmd.(CreateUrlShortenerCommand)
	if !ok {
		return fmt.Errorf("%w: %s", ErrUnexpectedCommand, cmd.Type())
	}
	return h.applicationService.Do(ctx, command)
}
