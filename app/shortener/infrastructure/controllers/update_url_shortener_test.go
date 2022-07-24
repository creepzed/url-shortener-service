package controllers

import (
	"context"
	"fmt"
	inmemoryBus "github.com/creepzed/url-shortener-service/app/shared/infrastructure/bus/inmemory"
	"github.com/creepzed/url-shortener-service/app/shortener/application/updating"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/mocks/storagemocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/controllers/request"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestUpdateUrlShortener(t *testing.T) {
	t.Parallel()
	t.Run("given a valid Url Short, It is update and returns 200", func(t *testing.T) {

		urlId := randomvalues.RandomUrlId()
		urlShortener := domain.RandomUrlShortener(urlId, vo.Enabled)
		auxEnabled := urlShortener.IsEnabled().Value()
		requestUpdate := request.RandomUpdateRequest(urlId, &auxEnabled).String()

		target := "/api/v1/shortener/"

		e := echoServer()
		req := httptest.NewRequest(http.MethodPatch, target, strings.NewReader(requestUpdate))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)
		ctx.SetPath(fmt.Sprintf("%s/:url_id", target))
		ctx.SetParamNames("url_id")
		ctx.SetParamValues(urlId)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(urlShortener, nil)
		mockRepository.
			On("Update", context.Background(), mock.AnythingOfType("domain.UrlShortener")).
			Return(nil)

		commandBusInMemory := inmemoryBus.NewCommandBusMemory()
		queryBusInMemory := inmemoryBus.NewQueryBusMemory()
		eventBusInMemory := inmemoryBus.NewEventBusInMemory()

		updateService := updating.NewUpdateApplicationService(mockRepository, eventBusInMemory)
		updateCommandHandler := updating.NewUpdateUrlShortenerCommandHandler(updateService)
		commandBusInMemory.Register(updating.UpdateUrlShortenerCommandType, updateCommandHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory, queryBusInMemory)
		err := urlShortenerController.Update(ctx)
		require.NoError(t, err)

		res := rec.Result()
		defer res.Body.Close()

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
		}
	})
	t.Run("given a invalid Url Short, It is not update and returns 400", func(t *testing.T) {

		urlId := randomvalues.InvalidUrlId()
		auxEnabled := vo.Enabled
		requestUpdate := request.RandomUpdateRequest(urlId, &auxEnabled).String()

		target := "/api/v1/shortener/"

		e := echoServer()
		req := httptest.NewRequest(http.MethodPatch, target, strings.NewReader(requestUpdate))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)
		ctx.SetPath(fmt.Sprintf("%s/:url_id", target))
		ctx.SetParamNames("url_id")
		ctx.SetParamValues(urlId)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)

		commandBusInMemory := inmemoryBus.NewCommandBusMemory()
		queryBusInMemory := inmemoryBus.NewQueryBusMemory()
		eventBusInMemory := inmemoryBus.NewEventBusInMemory()

		updateService := updating.NewUpdateApplicationService(mockRepository, eventBusInMemory)
		updateCommandHandler := updating.NewUpdateUrlShortenerCommandHandler(updateService)
		commandBusInMemory.Register(updating.UpdateUrlShortenerCommandType, updateCommandHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory, queryBusInMemory)
		err := urlShortenerController.Update(ctx)

		res := rec.Result()
		defer res.Body.Close()

		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.ErrorIs(t, err, exception.ErrInvalidUrlId)
		}
	})
	t.Run("given a valid Url Short, It is not found returns 404", func(t *testing.T) {

		urlId := randomvalues.RandomUrlId()
		auxEnabled := vo.Enabled
		requestUpdate := request.RandomUpdateRequest(urlId, &auxEnabled).String()

		target := "/api/v1/shortener/"

		e := echoServer()
		req := httptest.NewRequest(http.MethodPatch, target, strings.NewReader(requestUpdate))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)
		ctx.SetPath(fmt.Sprintf("%s/:url_id", target))
		ctx.SetParamNames("url_id")
		ctx.SetParamValues(urlId)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)
		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(domain.UrlShortener{}, nil)

		commandBusInMemory := inmemoryBus.NewCommandBusMemory()
		queryBusInMemory := inmemoryBus.NewQueryBusMemory()
		eventBusInMemory := inmemoryBus.NewEventBusInMemory()

		updateService := updating.NewUpdateApplicationService(mockRepository, eventBusInMemory)
		updateCommandHandler := updating.NewUpdateUrlShortenerCommandHandler(updateService)
		commandBusInMemory.Register(updating.UpdateUrlShortenerCommandType, updateCommandHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory, queryBusInMemory)
		err := urlShortenerController.Update(ctx)

		res := rec.Result()
		defer res.Body.Close()

		if assert.Error(t, err) {
			assert.Equal(t, http.StatusNotFound, rec.Code)
			assert.ErrorIs(t, err, exception.ErrUrlNotFound)
		}
	})
	t.Run("given a valid short url, when searching the database returns an error, it returns 500", func(t *testing.T) {

		urlId := randomvalues.RandomUrlId()
		auxEnabled := vo.Enabled
		requestUpdate := request.RandomUpdateRequest(urlId, &auxEnabled).String()

		target := "/api/v1/shortener/"

		e := echoServer()
		req := httptest.NewRequest(http.MethodPatch, target, strings.NewReader(requestUpdate))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)
		ctx.SetPath(fmt.Sprintf("%s/:url_id", target))
		ctx.SetParamNames("url_id")
		ctx.SetParamValues(urlId)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)

		mockRepository.
			On("FindById", context.Background(), mock.AnythingOfType("vo.UrlId")).
			Return(domain.UrlShortener{}, exception.ErrDataBase)

		commandBusInMemory := inmemoryBus.NewCommandBusMemory()
		queryBusInMemory := inmemoryBus.NewQueryBusMemory()
		eventBusInMemory := inmemoryBus.NewEventBusInMemory()

		updateService := updating.NewUpdateApplicationService(mockRepository, eventBusInMemory)
		updateCommandHandler := updating.NewUpdateUrlShortenerCommandHandler(updateService)
		commandBusInMemory.Register(updating.UpdateUrlShortenerCommandType, updateCommandHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory, queryBusInMemory)
		err := urlShortenerController.Update(ctx)

		res := rec.Result()
		defer res.Body.Close()

		if assert.Error(t, err) {
			assert.Equal(t, http.StatusInternalServerError, rec.Code)
			assert.ErrorIs(t, err, exception.ErrDataBase)
		}
	})
	t.Run("given a empty Url Short, It is not update and returns 400", func(t *testing.T) {

		urlId := ""
		auxEnabled := vo.Enabled
		requestUpdate := request.RandomUpdateRequest(urlId, &auxEnabled).String()

		target := "/api/v1/shortener/"

		e := echoServer()
		req := httptest.NewRequest(http.MethodPatch, target, strings.NewReader(requestUpdate))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)
		ctx.SetPath(fmt.Sprintf("%s/:url_id", target))
		ctx.SetParamNames("url_id")
		ctx.SetParamValues(urlId)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)

		commandBusInMemory := inmemoryBus.NewCommandBusMemory()
		queryBusInMemory := inmemoryBus.NewQueryBusMemory()
		eventBusInMemory := inmemoryBus.NewEventBusInMemory()

		updateService := updating.NewUpdateApplicationService(mockRepository, eventBusInMemory)
		updateCommandHandler := updating.NewUpdateUrlShortenerCommandHandler(updateService)
		commandBusInMemory.Register(updating.UpdateUrlShortenerCommandType, updateCommandHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory, queryBusInMemory)
		err := urlShortenerController.Update(ctx)

		res := rec.Result()
		defer res.Body.Close()

		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.ErrorIs(t, err, exception.ErrInvalidUrlId)
		}
	})
	t.Run("given a empty body Url Short, It is not update and returns 400", func(t *testing.T) {

		urlId := randomvalues.RandomUrlId()
		requestUpdate := ""

		target := "/api/v1/shortener/"

		e := echoServer()
		req := httptest.NewRequest(http.MethodPatch, target, strings.NewReader(requestUpdate))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)
		ctx.SetPath(fmt.Sprintf("%s/:url_id", target))
		ctx.SetParamNames("url_id")
		ctx.SetParamValues(urlId)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)

		commandBusInMemory := inmemoryBus.NewCommandBusMemory()
		queryBusInMemory := inmemoryBus.NewQueryBusMemory()
		eventBusInMemory := inmemoryBus.NewEventBusInMemory()

		updateService := updating.NewUpdateApplicationService(mockRepository, eventBusInMemory)
		updateCommandHandler := updating.NewUpdateUrlShortenerCommandHandler(updateService)
		commandBusInMemory.Register(updating.UpdateUrlShortenerCommandType, updateCommandHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory, queryBusInMemory)
		err := urlShortenerController.Update(ctx)

		res := rec.Result()
		defer res.Body.Close()

		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.ErrorIs(t, err, ErrInvalidRequestBody)
		}
	})

	t.Run("given a invalid body, It is not update and returns 400", func(t *testing.T) {

		urlId := randomvalues.RandomUrlId()
		requestUpdate := "&&&&&&"

		target := "/api/v1/shortener/"

		e := echoServer()
		req := httptest.NewRequest(http.MethodPatch, target, strings.NewReader(requestUpdate))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)
		ctx.SetPath(fmt.Sprintf("%s/:url_id", target))
		ctx.SetParamNames("url_id")
		ctx.SetParamValues(urlId)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)

		commandBusInMemory := inmemoryBus.NewCommandBusMemory()
		queryBusInMemory := inmemoryBus.NewQueryBusMemory()
		eventBusInMemory := inmemoryBus.NewEventBusInMemory()

		updateService := updating.NewUpdateApplicationService(mockRepository, eventBusInMemory)
		updateCommandHandler := updating.NewUpdateUrlShortenerCommandHandler(updateService)
		commandBusInMemory.Register(updating.UpdateUrlShortenerCommandType, updateCommandHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory, queryBusInMemory)
		err := urlShortenerController.Update(ctx)

		res := rec.Result()
		defer res.Body.Close()

		if assert.Error(t, err) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
			assert.ErrorIs(t, err, ErrInvalidRequestBody)
		}
	})
}
