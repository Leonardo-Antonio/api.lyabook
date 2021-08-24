package router

import (
	"github.com/Leonardo-Antonio/api.lyabook/src/handler"
	"github.com/Leonardo-Antonio/api.lyabook/src/middleware"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/labstack/echo/v4"
)

func Book(storageBook model.Ibook, storageUser model.IUser, app *echo.Echo) {
	book := handler.NewBook(storageBook, storageUser)

	group := app.Group(env.Data.BaseUrl + "/books")
	group.Use(middleware.Authorization().Admin)
	group.POST("/:format", book.Create)  // d -> digital, f -> fisico or df -> digital and fisico
	group.PUT("/:format/:id", book.Edit) // d -> digital, f -> fisico or df -> digital and fisico
	group.PATCH("/promotions/:id", book.AddPromotion)
}
