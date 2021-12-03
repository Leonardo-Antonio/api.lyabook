package router

import (
	"github.com/Leonardo-Antonio/api.lyabook/src/handler"
	"github.com/Leonardo-Antonio/api.lyabook/src/middleware"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/labstack/echo/v4"
)

func Category(storage model.ICategory, app *echo.Echo) {
	category := handler.NewCategory(storage)

	group := app.Group(env.Data.BaseUrl+"/categories", middleware.Authorization().Admin)
	group.POST("", category.AddMany)
	group.POST("/one", category.Add)
	group.DELETE("", category.DeleteById)
	group.PUT("/:id", category.Update)
}
