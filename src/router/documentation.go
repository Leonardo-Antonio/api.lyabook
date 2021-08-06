package router

import (
	"github.com/Leonardo-Antonio/api.lyabook/src/handler"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/labstack/echo/v4"
)

func Documentation(e *echo.Echo) {
	doc := handler.NewDocumentation()
	group := e.Group(env.Data.BaseUrl)
	group.GET("", doc.Index)
}
