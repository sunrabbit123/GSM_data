package data

import (
	"github.com/labstack/echo"

	"GSM_data/server/controller"
)

func RouteData(group *echo.Group) {
	group.GET("/trand_taste", controller.GetDataAboutTrandTaste)
}
