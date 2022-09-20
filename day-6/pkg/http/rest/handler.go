package rest

import (
	"github.com/labstack/echo/v4"
)

func InitHandlers() *echo.Echo {
	r := echo.New()

	r.GET("/api/hello", GetData)

	return r
}
