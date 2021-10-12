package router

import (
	"fmt"

	"github.com/Leonardo-Antonio/api.lyabook/src/handler"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/labstack/echo/v4"
)

func Report(storage model.Ibook, app *echo.Echo) {
	handler := handler.NewReport(storage)
	group := app.Group(fmt.Sprintf("%s/%s", env.Data.BaseUrl, "reports"))
	group.GET("/books", handler.AllBooks)
	group.GET("/books/stock/:stock", handler.ReportBooksByStock)
	group.GET("/data/books/stock/:stock", handler.SearchBookByStock)
	group.GET("/data/new/books/:limit", handler.NewBooks)
	group.GET("/new/books/:limit", handler.ReportNewBooks)
}
