package controllers

import (
	"errors"
	"github.com/creepzed/url-shortener-service/app/shared/application/command"
	"github.com/creepzed/url-shortener-service/app/shared/application/query"
	"github.com/labstack/echo/v4"
)

type urlShortenerController struct {
	commandBus command.CommandBus
	queryBus   query.QueryBus
}

var (
	ErrInvalidRequestBody = errors.New("the request body is invalid")
)

func NewUrlShortenerController(e *echo.Echo, commandBus command.CommandBus, queryBus query.QueryBus) *urlShortenerController {
	controller := &urlShortenerController{
		commandBus: commandBus,
		queryBus:   queryBus,
	}

	v1 := e.Group("/api/v1")
	{
		subGroup := v1.Group("/shortener")
		{
			subGroup.POST("", controller.Create)
			subGroup.GET("/:url_id", controller.Find)
			subGroup.PATCH("/:url_id", controller.Update)
		}
	}
	return controller

}
