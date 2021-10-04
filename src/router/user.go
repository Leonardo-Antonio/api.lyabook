package router

import (
	"github.com/Leonardo-Antonio/api.lyabook/src/handler"
	"github.com/Leonardo-Antonio/api.lyabook/src/middleware"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/labstack/echo/v4"
)

func User(storage model.IUser, app *echo.Echo) {
	user := handler.NewUser(storage)

	group := app.Group(env.Data.BaseUrl + "/users")
	group.PUT("/:id", user.Update)
	group.DELETE("/admin", user.DeleteById, middleware.Authorization().Manager)
	group.POST("/sign-up/:type", user.SignUp)
	group.POST("/log-in/:type", user.LogIn)
	group.POST("/verify", user.VerifyAccount)
	group.POST("/verification/id/:id/code/:code", user.VerifyAccountById)
	group.GET("/roles/admin", user.FindAllUsersByRol, middleware.Authorization().ManagerAndAdmin)
	group.GET("/count/admin", user.CountByRol)
	group.GET("/search/reniec/dni/:dni", user.SearchDni)
}
