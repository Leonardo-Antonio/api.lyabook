package handler

import (
	"net/http"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/labstack/echo/v4"
)

type book struct {
	storage model.Ibook
}

func NewBook(storage model.Ibook) *book {
	return &book{storage}
}

func (b *book) Create(ctx echo.Context) error {
	var book entity.Book
	if err := ctx.Bind(&book); err != nil {
		return response.New(ctx, http.StatusBadRequest, "la estructura no es valida", true, nil)
	}

}
