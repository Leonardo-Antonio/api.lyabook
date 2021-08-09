package router

import (
	"github.com/Leonardo-Antonio/api.lyabook/src/handler"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/labstack/echo/v4"
)

func Book(storage model.Ibook, app *echo.Echo) {
	book := handler.NewBook(storage)

	group := app.Group(env.Data.BaseUrl + "/books")
	group.POST("/:format", book.Create) // d -> digital, f -> fisico or df -> digital and fisico
}
