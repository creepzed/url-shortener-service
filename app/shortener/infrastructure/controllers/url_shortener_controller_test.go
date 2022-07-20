package controllers

import (
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/rest"
	"github.com/labstack/echo/v4"
)

func echoServer() *echo.Echo {
	e := echo.New()
	e.Validator = rest.NewValidator()
	return e
}
