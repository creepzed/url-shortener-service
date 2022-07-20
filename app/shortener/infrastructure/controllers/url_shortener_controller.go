package controllers

import (
	"github.com/creepzed/url-shortener-service/app/shared/application/command"
	"github.com/creepzed/url-shortener-service/app/shared/application/query"
	"github.com/labstack/echo/v4"
)

type createUrlShortenerController struct {
	commandBus command.CommandBus
	queryBus   query.QueryBus
}

func NewUrlShortenerController(e *echo.Echo, commandBus command.CommandBus, queryBus query.QueryBus) *createUrlShortenerController {
	controller := &createUrlShortenerController{
		commandBus: commandBus,
		queryBus:   queryBus,
	}

	v1 := e.Group("/api/v1")
	{
		subGroup := v1.Group("/shortener")
		{
			subGroup.POST("", controller.Create)
			subGroup.GET("/:url_id", controller.Find)
		}
	}
	return controller

}
