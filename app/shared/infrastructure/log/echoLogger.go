package log

import (
	"github.com/labstack/echo/v4"
	"time"
)

func EchoLogger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()

			err := next(c)
			if err != nil {
				c.Error(err)
			}

			req := c.Request()
			res := c.Response()

			id := req.Header.Get(echo.HeaderXRequestID)
			if id == "" {
				id = res.Header().Get(echo.HeaderXRequestID)
			}

			logger := Logger()

			logger.
				WithField("status", res.Status).
				WithField("latency", time.Since(start).String()).
				WithField("id", id).
				WithField("method", req.Method).
				WithField("uri", req.RequestURI).
				WithField("host", req.Host).
				WithField("remote_ip", c.RealIP())

			n := res.Status

			switch {
			case n >= 500:
				logger.Error("rest error")
			case n >= 400:
				logger.Warn("client error")
			case n >= 300:
				logger.Info("redirection")
			default:
				logger.Info("success")

			}
			return nil
		}
	}
}
