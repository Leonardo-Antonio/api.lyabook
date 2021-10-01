package router

import (
	"fmt"

	"github.com/Leonardo-Antonio/api.lyabook/src/handler"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/labstack/echo/v4"
)

func Report(app *echo.Echo) {
	handler := handler.NewReport()
	group := app.Group(fmt.Sprintf("%s/%s", env.Data.BaseUrl, "reports"))
	group.GET("", handler.Create)
}
