package controllers

import (
	"errors"
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/infrastructure/log"
	"github.com/creepzed/url-shortener-service/app/shortener/application/creating"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/controllers/request"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/controllers/response"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Create Url Short godoc
// @Summary      Add an Url Short
// @Description  add by json Url Short
// @Tags         shortener
// @Accept       json
// @Produce      json
// @Param        shortener body     request.UrlShortenerRequestCreate  true  "Add Url"
// @Success      200      {object}  response.OutputResponse
// @Failure      400      {object}  map[string]interface{}
// @Failure      404      {object}  map[string]interface{}
// @Failure      500      {object}  map[string]interface{}
// @Router       /api/v1/shortener  [post]
func (ctrl *urlShortenerController) Create(c echo.Context) (err error) {
	request := new(request.UrlShortenerRequestCreate)

	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		return fmt.Errorf("%w: %s", ErrInvalidRequestBody, err.Error())
	}

	if err = c.Validate(request); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		return err
	}

	urlId, err := ctrl.kgs.GetKey()
	if err != nil {
		log.Fatal("error fatal ksg: %s", err.Error())
	}

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
	resp := response.OutputResponse{
		UrlId:       urlId,
		IsEnabled:   vo.Enabled,
		OriginalUrl: request.OriginalUrl,
		UserId:      request.UserId,
	}
	return c.JSON(http.StatusCreated, resp)
}
