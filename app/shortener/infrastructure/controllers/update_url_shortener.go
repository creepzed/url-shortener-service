package controllers

import (
	"errors"
	"fmt"
	"github.com/creepzed/url-shortener-service/app/shared/domain/vo"
	"github.com/creepzed/url-shortener-service/app/shortener/application/updating"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/creepzed/url-shortener-service/app/shortener/infrastructure/controllers/request"
	"github.com/labstack/echo/v4"
	"net/http"
)

// Update godoc
// @Summary      Update an Url Short
// @Description  Update by json Url Short
// @Tags         shortener
// @Accept       json
// @Produce      json
// @Param        url_Id          path      string                true  "Url Id"
// @Param        shortener       body      request.UpdateRequest true  "Update Url Short"
// @Failure      400             {object}  map[string]interface{}
// @Failure      404             {object}  map[string]interface{}
// @Failure      500             {object}  map[string]interface{}
// @Router    	 /api/v1/shortener/{url_id}  [patch]

func (ctrl *urlShortenerController) Update(c echo.Context) (err error) {
	urlId := c.Param("url_id")

	if len(urlId) == 0 {
		c.JSON(http.StatusBadRequest, echo.Map{"message": "no parameter in request"})
		return exception.ErrInvalidUrlId
	}

	if err := vo.IsValidUrlId(urlId); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		return err
	}

	request := new(request.UpdateRequest)

	if err := c.Bind(request); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		return fmt.Errorf("%w: %s", ErrInvalidRequestBody, err.Error())
	}

	if request.IsEnabled == nil && request.OriginalUrl == "" {
		c.JSON(http.StatusBadRequest, echo.Map{"message": "request body cannot be empty"})
		return ErrInvalidRequestBody
	}

	cmd := updating.NewUpdateUrlShortenerCommand(urlId, request.IsEnabled, request.OriginalUrl)

	err = ctrl.commandBus.Dispatch(c.Request().Context(), cmd)
	if err != nil {
		codeErr := http.StatusInternalServerError
		switch {
		case errors.Is(err, exception.ErrInvalidUrlId),
			errors.Is(err, exception.ErrInvalidOriginalUrl),
			errors.Is(err, exception.ErrInvalidUserId),
			errors.Is(err, exception.ErrEmptyUrlId),
			errors.Is(err, exception.ErrEmptyOriginalUrl):
			codeErr = http.StatusBadRequest
		case errors.Is(err, exception.ErrUrlNotFound):
			codeErr = http.StatusNotFound
		}
		c.JSON(codeErr, echo.Map{"message": err.Error()})
		return err
	}
	return c.NoContent(http.StatusOK)
}
