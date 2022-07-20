package creating

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/application/command"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/bus/busmocks/commandmocks"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/bus/busmocks/eventmocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/storage/storagemocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCreateUrlShortenerCommandHandler(t *testing.T) {
	t.Parallel()
	t.Run("given a valid registered command it is executed", func(t *testing.T) {

		//var commandMockType command.Type = "command.mock"
		//cmdMock := commandmocks.NewCommand(t)
		//cmdMock.On("Type").Return(commandMockType)

		cmd := NewCreateUrlShortenerCommand(randomvalues.RandomUrlId(), randomvalues.RandomOriginalUrl(), randomvalues.RandomUserId())

		repositoryMock := storagemocks.NewUrlShortenerRepository(t)

		repositoryMock.
			On("Create", context.Background(), mock.AnythingOfType("domain.UrlShortener")).
			Return(nil)

		repositoryMock.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(domain.UrlShortener{}, nil)

		eventBusMock := eventmocks.NewEventBus(t)
		eventBusMock.
			On("Publish", context.Background(), mock.AnythingOfType("[]event.Event")).
			Return(nil)

		service := NewCreateApplicationService(repositoryMock, eventBusMock)

		handler := NewCreateUrlShortenerCommandHandler(service)
		err := handler.Handle(context.Background(), cmd)

		assert.NoError(t, err)
	})

	t.Run("given a valid unregistered command, return an error", func(t *testing.T) {

		var commandMockType command.Type = "command.mock"
		cmdMock := commandmocks.NewCommand(t)
		cmdMock.On("Type").Return(commandMockType)

		repositoryMock := storagemocks.NewUrlShortenerRepository(t)
		eventBusMock := eventmocks.NewEventBus(t)
		service := NewCreateApplicationService(repositoryMock, eventBusMock)

		handler := NewCreateUrlShortenerCommandHandler(service)
		err := handler.Handle(context.Background(), cmdMock)

		require.Error(t, err)
		assert.ErrorIs(t, err, ErrUnexpectedCommand)
	})
}
