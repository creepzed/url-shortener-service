package main

import (
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/rest"
	"time"
)

const (
	host      = ""
	port      = 8080
	dbTimeOut = 5 * time.Second
)

func main() {
	echo := rest.New()

	echo.Logger.Fatal(echo.StartServer(rest.Setup(host, port)))
}
