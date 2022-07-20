package controllers

import (
	"errors"
	"github.com/creepzed/url-shortener-service/app/shortener/application/finding"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl *createUrlShortenerController) Find(c echo.Context) (err error) {
	urlId := c.Param("url_id")

	if len(urlId) == 0 {
		c.JSON(http.StatusBadRequest, echo.Map{"message": "no parameter in request"})
	}

	if err := vo.IsValidUrlId(urlId); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
	}

	qry := finding.NewFindUrlShortenerQuery(urlId)

	result, err := ctrl.queryBus.Execute(c.Request().Context(), qry)
	if err != nil {
		codeErr := http.StatusInternalServerError
		switch {
		case errors.Is(err, exception.ErrInvalidUrlId):
			codeErr = http.StatusBadRequest
		}
		c.JSON(codeErr, echo.Map{"message": err.Error()})
		return err
	}
	return c.JSON(http.StatusOK, result)
}