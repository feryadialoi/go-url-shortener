package respond

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type Response struct {
	Data   interface{} `json:"data"`
	Errors interface{} `json:"errors"`
}

func ResponseInternalServerError(c echo.Context, err error) error {
	return echo.NewHTTPError(http.StatusInternalServerError, Response{
		Errors: err.Error(),
	})
}

func ResponseBadRequest(c echo.Context, err error) error {
	return echo.NewHTTPError(http.StatusBadRequest, Response{
		Errors: err.Error(),
	})
}

func ResponseNotFound(c echo.Context, err error) error {
	return echo.NewHTTPError(http.StatusNotFound, Response{
		Errors: err.Error(),
	})
}
