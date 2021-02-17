package router

import (
	"fmt"
	"net/http"

	"github.com/Goolgae/GSM_data/server/router/data"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Router(debug bool) *echo.Echo {
	e := echo.New()
	e.Debug = debug

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	api := e.Group("/api")

	v1 := api.Group("/v1")

	dataApi := v1.Group("/data")

	data.RouteData(dataApi)

	return e
}

func Test() {
	fmt.Println("Test")
	return
}
