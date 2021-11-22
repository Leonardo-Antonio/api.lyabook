package router

import (
	"fmt"

	"github.com/Leonardo-Antonio/api.lyabook/src/handler"
	"github.com/Leonardo-Antonio/api.lyabook/src/middleware"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/labstack/echo/v4"
)

func ComplaintsBook(storage model.IComplaintsBook, app *echo.Echo) {
	complaintsBook := handler.NewComplaintsBook(storage)
	group := app.Group(fmt.Sprintf("%s/%s", env.Data.BaseUrl, "claims"))
	group.POST("", complaintsBook.Add, middleware.Authorization().Client)
	group.GET("", complaintsBook.GetAll, middleware.Authorization().Admin)
	group.GET("/amount", complaintsBook.CountClaims)
	group.POST("/response/:id", complaintsBook.ResponseClaim, middleware.Authorization().Admin)
}
