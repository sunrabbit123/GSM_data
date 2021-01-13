package router

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Router(debug bool) *echo.Echo {
	e := echo.New()
	e.Debug = debug

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	apiV1 := e.Group("/api/v1")
	apiV1.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	return e
}
