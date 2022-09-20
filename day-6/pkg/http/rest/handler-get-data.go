package rest

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetData(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{
		"msg": "hello",
	})
}
