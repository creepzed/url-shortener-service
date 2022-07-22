package controllers

import (
	"context"
	"fmt"
	inmemoryBus "github.com/creepzed/url-shortener-service/app/shared/infrastructure/bus/inmemory"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/utils"
	"github.com/creepzed/url-shortener-service/app/shortener/application/finding"
	"github.com/creepzed/url-shortener-service/app/shortener/domain"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/mocks/storagemocks"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/controllers/transformer"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestFindUrlShortener(t *testing.T) {
	t.Parallel()
	t.Run("given a valid Url Short, return data and code 200", func(t *testing.T) {

		urlId := randomvalues.RandomUrlId()
		urlShortener := domain.RandomUrlShortener(urlId, vo.Enabled)

		target := "/api/v1/shortener/"

		e := echoServer()
		req := httptest.NewRequest(http.MethodGet, target, nil)
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

		commandBusInMemory := inmemoryBus.NewCommandBusMemory()
		queryBusInMemory := inmemoryBus.NewQueryBusMemory()

		transform := transformer.NewTransformer()

		findService := finding.NewFindApplicationService(mockRepository, transform)
		findQueryHandler := finding.NewFindUrlShortenerQueryHandler(findService)
		queryBusInMemory.Register(finding.FindUrlShortenerQueryType, findQueryHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory, queryBusInMemory)
		err := urlShortenerController.Find(ctx)

		res := rec.Result()
		defer res.Body.Close()

		responseExpected, err := transform.Transform(urlShortener)
		require.NoError(t, err)
		jsonExpected := fmt.Sprintf("%s\n", utils.EntityToJson(responseExpected))

		if assert.NoError(t, err) {
			assert.Equal(t, http.StatusOK, rec.Code)
			assert.Equal(t, jsonExpected, rec.Body.String())
		}
	})

	t.Run("given a valid Url Short, return error data and code 404", func(t *testing.T) {

		urlId := randomvalues.RandomUrlId()

		target := "/api/v1/shortener/"

		e := echoServer()
		req := httptest.NewRequest(http.MethodGet, target, nil)
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

		transform := transformer.NewTransformer()

		findService := finding.NewFindApplicationService(mockRepository, transform)

		findQueryHandler := finding.NewFindUrlShortenerQueryHandler(findService)
		queryBusInMemory.Register(finding.FindUrlShortenerQueryType, findQueryHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory, queryBusInMemory)
		err := urlShortenerController.Find(ctx)

		res := rec.Result()
		defer res.Body.Close()

		if assert.Error(t, err) {
			assert.ErrorIs(t, err, exception.ErrUrlNotFound)
			assert.Equal(t, http.StatusNotFound, rec.Code)
		}
	})

	t.Run("given a invalid Url Short, return error data and code 400", func(t *testing.T) {

		urlId := randomvalues.InvalidUrlId()

		target := "/api/v1/shortener/"

		e := echoServer()
		req := httptest.NewRequest(http.MethodGet, target, nil)
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()

		ctx := e.NewContext(req, rec)
		ctx.SetPath(fmt.Sprintf("%s/:url_id", target))
		ctx.SetParamNames("url_id")
		ctx.SetParamValues(urlId)

		mockRepository := storagemocks.NewUrlShortenerRepository(t)

		commandBusInMemory := inmemoryBus.NewCommandBusMemory()
		queryBusInMemory := inmemoryBus.NewQueryBusMemory()

		transform := transformer.NewTransformer()

		findService := finding.NewFindApplicationService(mockRepository, transform)

		findQueryHandler := finding.NewFindUrlShortenerQueryHandler(findService)
		queryBusInMemory.Register(finding.FindUrlShortenerQueryType, findQueryHandler)

		urlShortenerController := NewUrlShortenerController(e, commandBusInMemory, queryBusInMemory)
		err := urlShortenerController.Find(ctx)

		res := rec.Result()
		defer res.Body.Close()

		if assert.Error(t, err) {
			assert.ErrorIs(t, err, exception.ErrInvalidUrlId)
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}
	})
}
