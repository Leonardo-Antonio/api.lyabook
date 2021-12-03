package router

import (
	"github.com/Leonardo-Antonio/api.lyabook/src/handler"
	"github.com/labstack/echo/v4"
)

func Documentation(e *echo.Echo) {
	doc := handler.NewDocumentation()
	e.GET("/", doc.Index)
}
