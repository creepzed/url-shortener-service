package controllers

import (
	"context"
	inmemoryBus "github.com/creepzed/url-shortener-service/app/shared/infrastructure/bus/inmemory"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/rest"
	"github.com/creepzed/url-shortener-service/app/shortener/application/creating"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/mother"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/controllers/request"
	inmemoryDB "github.com/creepzed/url-shortener-service/app/shortener/infrastructure/storage/inmemory"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/storage/storagemocks"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestCreateUrlShortenerController(t *testing.T) {
	t.Parallel()

	t.Run("given a valid request it returns 201", func(t *testing.T) {
		target := "/api/v1/shortener"
		requestString := request.RandomUrlShortenerCreateRequest().String()

		e := echoServer()
		req := httptest.NewRequest(http.MethodPost, target, strings.NewReader(requestString))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)

		dataBaseInMemory := inmemoryDB.NewDataBaseInMemory()
		eventBusInMemory := inmemoryBus.NewEventBusInMemory()
		commandBusInMemory := inmemoryBus.NewCommandBusMemory()

		createService := creating.NewCreateApplicationService(dataBaseInMemory, eventBusInMemory)

		createCommandHandler := creating.NewCreateUrlShortenerCommandHandler(createService)

		commandBusInMemory.Register(creating.CreateUrlShortenerCommandType, createCommandHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory)
		err := urlShortenerController.Create(ctx)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}
	})

	t.Run("given an invalid request when the origin url is wrong, then returns 400 with an error message", func(t *testing.T) {
		target := "/api/v1/shortener"
		requestString := request.FailRequestWithWrongOriginalUrl().String()

		e := echoServer()
		req := httptest.NewRequest(http.MethodPost, target, strings.NewReader(requestString))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)

		dataBaseInMemory := inmemoryDB.NewDataBaseInMemory()
		eventBusInMemory := inmemoryBus.NewEventBusInMemory()
		commandBusInMemory := inmemoryBus.NewCommandBusMemory()

		createService := creating.NewCreateApplicationService(dataBaseInMemory, eventBusInMemory)

		createCommandHandler := creating.NewCreateUrlShortenerCommandHandler(createService)

		commandBusInMemory.Register(creating.CreateUrlShortenerCommandType, createCommandHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory)
		err := urlShortenerController.Create(ctx)

		res := rec.Result()
		defer res.Body.Close()

		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		}
	})

	t.Run("given an invalid request when the user id is wrong, then returns 400 with an error message", func(t *testing.T) {
		target := "/api/v1/shortener"
		requestString := request.FailRequestWithWrongUserId().String()

		e := echoServer()
		req := httptest.NewRequest(http.MethodPost, target, strings.NewReader(requestString))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)

		dataBaseInMemory := inmemoryDB.NewDataBaseInMemory()
		eventBusInMemory := inmemoryBus.NewEventBusInMemory()
		commandBusInMemory := inmemoryBus.NewCommandBusMemory()

		createService := creating.NewCreateApplicationService(dataBaseInMemory, eventBusInMemory)

		createCommandHandler := creating.NewCreateUrlShortenerCommandHandler(createService)

		commandBusInMemory.Register(creating.CreateUrlShortenerCommandType, createCommandHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory)
		err := urlShortenerController.Create(ctx)

		res := rec.Result()
		defer res.Body.Close()

		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		}
	})

	t.Run("given a invalid request it returns 400", func(t *testing.T) {
		target := "/api/v1/shortener"
		requestString := "{}{"
		e := echoServer()
		req := httptest.NewRequest(http.MethodPost, target, strings.NewReader(requestString))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)

		dataBaseInMemory := inmemoryDB.NewDataBaseInMemory()
		eventBusInMemory := inmemoryBus.NewEventBusInMemory()
		commandBusInMemory := inmemoryBus.NewCommandBusMemory()

		createService := creating.NewCreateApplicationService(dataBaseInMemory, eventBusInMemory)

		createCommandHandler := creating.NewCreateUrlShortenerCommandHandler(createService)

		commandBusInMemory.Register(creating.CreateUrlShortenerCommandType, createCommandHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory)
		err := urlShortenerController.Create(ctx)

		res := rec.Result()
		defer res.Body.Close()

		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, res.StatusCode)
		}
	})

	t.Run("given the duplicate UrlId returns 409", func(t *testing.T) {
		target := "/api/v1/shortener"

		urlId := randomvalues.RandomUrlId()
		originalUrl := randomvalues.RandomOriginalUrl()
		userId := randomvalues.RandomUserId()

		requestString := request.RandomUrlShortenerCreateRequest().String()

		e := echoServer()
		req := httptest.NewRequest(http.MethodPost, target, strings.NewReader(requestString))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)

		existUrl := domain.NewUrlShortener(mother.UrlId(urlId), mother.OriginalUrl(originalUrl), mother.UserId(userId))

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(existUrl, nil)

		eventBusInMemory := inmemoryBus.NewEventBusInMemory()
		commandBusInMemory := inmemoryBus.NewCommandBusMemory()

		createService := creating.NewCreateApplicationService(mockRepository, eventBusInMemory)

		createCommandHandler := creating.NewCreateUrlShortenerCommandHandler(createService)

		commandBusInMemory.Register(creating.CreateUrlShortenerCommandType, createCommandHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory)
		err := urlShortenerController.Create(ctx)

		res := rec.Result()
		defer res.Body.Close()

		if assert.Error(t, err) {
			assert.Equal(t, http.StatusConflict, res.StatusCode)
		}
	})
}

func echoServer() *echo.Echo {
	e := echo.New()
	e.Validator = rest.NewValidator()
	return e
}
