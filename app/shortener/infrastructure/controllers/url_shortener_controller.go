package controllers

import (
	"errors"
	"github.com/creepzed/url-shortener-service/app/key-generation-service/application/generating"
	"github.com/creepzed/url-shortener-service/app/shared/application/command"
	"github.com/creepzed/url-shortener-service/app/shared/application/query"
	"github.com/labstack/echo/v4"
)

type urlShortenerController struct {
	commandBus command.CommandBus
	queryBus   query.QueryBus
	kgs        generating.KeyGenerateService
}

var (
	ErrInvalidRequestBody = errors.New("the request body is invalid")
)

func NewUrlShortenerController(e *echo.Echo, commandBus command.CommandBus, queryBus query.QueryBus, kgs generating.KeyGenerateService) *urlShortenerController {
	controller := &urlShortenerController{
		commandBus: commandBus,
		queryBus:   queryBus,
		kgs:        kgs,
	}

	v1 := e.Group("/api/v1")
	{
		subGroup := v1.Group("/shortener")
		{
			subGroup.POST("", controller.Create)
			subGroup.GET("/:url_id", controller.Find)
			subGroup.PATCH("/:url_id", controller.Update)
			subGroup.GET("/user/:usr_id", controller.GetAll)
		}
	}

	return controller

}
