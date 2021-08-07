package router

import (
	"github.com/Leonardo-Antonio/api.lyabook/src/handler"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/labstack/echo/v4"
)

func User(storage model.IUser, app *echo.Echo) {
	user := handler.NewUser(storage)

	group := app.Group(env.Data.BaseUrl + "/users")
	group.POST("/sign-up/:type", user.SignUp)
}
