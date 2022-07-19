package rest

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"os"
	"time"
)

var startTime time.Time

func init() {
	startTime = time.Now()
}

type Health struct {
	Status      string `json:"status"`       //status
	ServiceName string `json:"service_name"` //service name
	Version     string `json:"version"`      //version
	Uptime      string `json:"uptime"`       //uptime
	Environment string `json:"environment"`  //environment
	Region      string `json:"region"`       //region
	CloudName   string `json:"cloud_name"`   //cloud name
}

type healthHandler struct {
}

func NewHealthHandler(e *echo.Echo) {
	h := &healthHandler{}
	e.GET("/health", h.HealthCheck)
}

func (p *healthHandler) HealthCheck(c echo.Context) error {
	versionApp := os.Getenv("VERSION_APP")
	healthCheck := Health{
		Status:      "UP",
		ServiceName: os.Getenv("SERVICE_NAME"),
		Version:     versionApp,
		Uptime:      time.Since(startTime).String(),
		Environment: os.Getenv("ENVIRONMENT"),
		Region:      os.Getenv("REGION"),
		CloudName:   os.Getenv("CLOUD_NAME"),
	}
	return c.JSON(http.StatusOK, healthCheck)
}
