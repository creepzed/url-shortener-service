package updating

import (
	"context"
	"errors"
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/application/command"
)

type UpdateUrlShortenerCommandHandler struct {
	applicationService UpdateApplicationService
}

var (
	ErrUnexpectedCommand = errors.New("unexpected command")
)

func NewUpdateUrlShortenerCommandHandler(applicationService UpdateApplicationService) UpdateUrlShortenerCommandHandler {
	return UpdateUrlShortenerCommandHandler{
		applicationService: applicationService,
	}
}

func (h UpdateUrlShortenerCommandHandler) Handle(ctx context.Context, cmd command.Command) error {
	command, ok := cmd.(UpdateUrlShortenerCommand)
	if !ok {
		return fmt.Errorf("%w: %s", ErrUnexpectedCommand, cmd.Type())
	}
	return h.applicationService.Do(ctx, command)
}
