package data

import (
	"net/http"

	"github.com/labstack/echo"

	"GSM_data/server/controller"
)

func RouteData(group *echo.Group) {
	group.GET("/trand_taste", controller.GetDataAboutTrandTaste)
	group.GET("", func(c echo.Context) error {
		return c.String(http.StatusOK, "api/v1/data")
	})
}
