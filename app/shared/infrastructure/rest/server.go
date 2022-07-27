package rest

import (
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
	_ "github.com/creepzed/url-shortener-service/docs/openapi"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
	"net/http"
	"time"
)

const (
	httpReadTimeout  = 3 * time.Minute
	httpWriteTimeout = 3 * time.Minute
)

// @title Url Shortener Service
// @version 0.1
// @description This is a service.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url https://github.com/creepzed/url-shortener-service
// @contact.email rodrigo.cuevas.morales@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

func New() *echo.Echo {

	echo := echo.New()

	echo.Use(log.EchoLogger())
	echo.Use(echoMiddleware.Logger())
	echo.Use(echoMiddleware.Recover())
	echo.Use(echoMiddleware.CORS())

	echo.Validator = NewValidator()

	echo.HideBanner = true

	NewHealthHandler(echo)
	echo.GET("/swagger/*", echoSwagger.WrapHandler)
	return echo
}

func Setup(host string, port string) *http.Server {
	server := &http.Server{
		Addr:         fmt.Sprintf("%s:%s", host, port),
		ReadTimeout:  httpReadTimeout,
		WriteTimeout: httpWriteTimeout,
	}
	return server

}
