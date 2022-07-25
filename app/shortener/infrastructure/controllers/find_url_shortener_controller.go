package controllers

import (
	"errors"
	"github.com/creepzed/url-shortener-service/app/shortener/application/finding"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl *urlShortenerController) Find(c echo.Context) (err error) {
	urlId := c.Param("url_id")

	if len(urlId) == 0 {
		c.JSON(http.StatusBadRequest, echo.Map{"message": "no parameter in request"})
		return exception.ErrInvalidUrlId
	}

	if err := vo.IsValidUrlId(urlId); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		return err
	}

	metadata := finding.Metadata{
		"Header":    c.Request().Header,
		"Method":    c.Request().Method,
		"Uri":       c.Request().RequestURI,
		"Host":      c.Request().Host,
		"Remote_ip": c.RealIP(),
	}

	qry := finding.NewFindUrlShortenerQuery(urlId, metadata)

	result, err := ctrl.queryBus.Execute(c.Request().Context(), qry)
	if err != nil {
		codeErr := http.StatusInternalServerError

		switch {
		case errors.Is(err, exception.ErrInvalidUrlId):
			codeErr = http.StatusBadRequest
		case errors.Is(err, exception.ErrUrlNotFound):
			codeErr = http.StatusNotFound
		}
		c.JSON(codeErr, echo.Map{"message": err.Error()})
		return err
	}
	return c.JSON(http.StatusOK, result)
}
