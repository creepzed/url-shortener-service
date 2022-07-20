package controllers

import (
	"errors"
	"github.com/creepzed/url-shortener-service/app/shared/application/command"
	"github.com/creepzed/url-shortener-service/app/shortener/application/creating"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo/randomvalues"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/controllers/request"
	"github.com/labstack/echo/v4"
	"net/http"
)

type createUrlShortenerController struct {
	commandBus command.CommandBus
}

func NewUrlShortenerController(e *echo.Echo, commandBus command.CommandBus) *createUrlShortenerController {
	controller := &createUrlShortenerController{
		commandBus: commandBus,
	}

	v1 := e.Group("/api/v1")
	{
		subGroup := v1.Group("/shortener")
		{
			subGroup.POST("", controller.Create)
		}
	}
	return controller

}

func (ctrl *createUrlShortenerController) Create(c echo.Context) (err error) {
	request := new(request.UrlShortenerCreateRequest)

	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		return err
	}

	if err = c.Validate(request); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		return err
	}

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
		case errors.Is(err, exception.ErrInvalidOriginalUrl),
			errors.Is(err, exception.ErrInvalidOriginalUrl),
			errors.Is(err, exception.ErrInvalidUserId):
			codeErr = http.StatusBadRequest
		}
		c.JSON(codeErr, echo.Map{"message": err.Error()})
		return err
	}

	return c.NoContent(http.StatusCreated)
}
