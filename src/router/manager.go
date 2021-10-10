package router

import (
	"fmt"

	"github.com/Leonardo-Antonio/api.lyabook/src/handler"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/labstack/echo/v4"
)

func Manager(app *echo.Echo) {
	manager := handler.NewManager()
	group := app.Group(fmt.Sprintf("%s/managers", env.Data.BaseUrl))
	group.POST("/administrators/message", manager.MessageAdmin)
}
