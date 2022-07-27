package updating

import (
	"context"
	"github.com/creepzed/url-shortener-service/app/shared/application/command"
	"github.com/creepzed/url-shortener-service/app/shared/application/mocks/commandmocks"
	"github.com/creepzed/url-shortener-service/app/shared/application/mocks/eventmocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/mocks/storagemocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestUpdateUrlShortenerCommandHandler(t *testing.T) {
	t.Parallel()
	t.Run("given a valid update command it is executed", func(t *testing.T) {

		urlId := randomvalues.RandomUrlId()
		urlExpected := domain.RandomUrlShortener(urlId, vo.Enabled)

		auxIsEnabled := urlExpected.IsEnabled().Value()
		cmd := NewUpdateUrlShortenerCommand(urlId, &auxIsEnabled, urlExpected.OriginalUrl().Value())

		mockRepository := storagemocks.NewUrlShortenerRepository(t)

		mockRepository.
			On("Update", context.Background(), mock.AnythingOfType("domain.UrlShortener")).
			Return(nil)

		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(urlExpected, nil)

		eventBusMock := eventmocks.NewEventBus(t)
		eventBusMock.
			On("Publish", context.Background(), mock.AnythingOfType("[]event.Event")).
			Return(nil)

		service := NewUpdateApplicationService(mockRepository, eventBusMock)

		handler := NewUpdateUrlShortenerCommandHandler(service)
		err := handler.Handle(context.Background(), cmd)

		assert.NoError(t, err)
	})

	t.Run("given a valid unregistered command, return an error", func(t *testing.T) {

		var commandMockType command.Type = "command.mock"
		cmdMock := commandmocks.NewCommand(t)
		cmdMock.On("Type").Return(commandMockType)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		eventBusMock := eventmocks.NewEventBus(t)
		service := NewUpdateApplicationService(mockRepository, eventBusMock)

		handler := NewUpdateUrlShortenerCommandHandler(service)
		err := handler.Handle(context.Background(), cmdMock)

		require.Error(t, err)
		assert.ErrorIs(t, err, ErrUnexpectedCommand)
	})
}
