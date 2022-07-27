package controllers

import (
	"errors"
	"github.com/creepzed/url-shortener-service/app/shortener/application/getting"
	"github.com/creepzed/url-shortener-service/app/shortener/domain/vo"

	"github.com/creepzed/url-shortener-service/app/shortener/domain/exception"
	"github.com/labstack/echo/v4"
	"net/http"
)

// ListUrlShort by User Id godoc
// @Summary      List Url Shorts
// @Description  get Url Shorts by User Id
// @Tags         shortener
// @Accept       json
// @Produce      json
// @Param        user_id path   string  true  "User Id"
// @Success      200  {array}   response.OutputResponse
// @Failure      400  {object}  map[string]interface{}
// @Failure      404  {object}  map[string]interface{}
// @Failure      500  {object}  map[string]interface{}
// @Router      /api/v1/shortener/user/{user_id} [get]

func (ctrl *urlShortenerController) GetAll(c echo.Context) (err error) {
	userId := c.Param("usr_id")

	if len(userId) == 0 {
		c.JSON(http.StatusBadRequest, echo.Map{"message": "no parameter in request"})
		return exception.ErrInvalidUserId
	}

	if err := vo.IsValidUserId(userId); err != nil {
		c.JSON(http.StatusBadRequest, echo.Map{"message": err.Error()})
		return err
	}

	qry := getting.NewGetAllUrlShortenerQuery(userId)

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
