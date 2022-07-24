package controllers

import (
	"errors"
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shortener/application/creating"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/controllers/request"
	"github.com/labstack/echo/v4"
	"net/http"
)

func (ctrl *urlShortenerController) Create(c echo.Context) (err error) {
	request := new(request.UrlShortenerRequestCreate)

	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		return fmt.Errorf("%w: %s", ErrInvalidBodyRequest, err.Error())
	}

	if err = c.Validate(request); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		return err
	}

	//TODO: I need to work at Key Generator Service
	urlId := randomvalues.RandomUrlId()

	cmd := creating.NewCreateUrlShortenerCommand(
		urlId,
		request.OriginalUrl,
		request.UserId,
	)

	err = ctrl.commandBus.Dispatch(c.Request().Context(), cmd)
	if err != nil {
		codeErr := http.StatusInternalServerError

		switch {
		case errors.Is(err, exception.ErrUrlIdDuplicate):
			codeErr = http.StatusConflict
		case errors.Is(err, exception.ErrInvalidUrlId),
			errors.Is(err, exception.ErrInvalidOriginalUrl),
			errors.Is(err, exception.ErrInvalidUserId),
			errors.Is(err, exception.ErrEmptyUrlId),
			errors.Is(err, exception.ErrEmptyOriginalUrl):
			codeErr = http.StatusBadRequest
		}
		c.JSON(codeErr, echo.Map{"message": err.Error()})
		return err
	}

	return c.NoContent(http.StatusCreated)
}
