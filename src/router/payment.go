package router

import (
	"fmt"

	"github.com/Leonardo-Antonio/api.lyabook/src/handler"
	"github.com/Leonardo-Antonio/api.lyabook/src/middleware"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/labstack/echo/v4"
)

func Payment(storage model.IPayment, app *echo.Echo) {
	handler := handler.NewPayment(storage)
	group := app.Group(fmt.Sprintf("%s/%s", env.Data.BaseUrl, "payments"))
	group.GET("/boleta/:id", handler.GetById)
	group.GET("/reports/books/sold", handler.GetAllBooksSold, middleware.Authorization().Manager)
}
