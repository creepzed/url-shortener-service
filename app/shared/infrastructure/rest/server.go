package rest

import (
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
	"github.com/labstack/echo/v4"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"net/http"
	"time"
)

const (
	httpReadTimeout  = 3 * time.Minute
	httpWriteTimeout = 3 * time.Minute
)

func New() *echo.Echo {

	echo := echo.New()

	echo.Use(log.EchoLogger())
	echo.Use(echoMiddleware.Logger())
	echo.Use(echoMiddleware.Recover())
	echo.Use(echoMiddleware.CORS())

	echo.Validator = NewValidator()

	echo.HideBanner = true

	NewHealthHandler(echo)

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
