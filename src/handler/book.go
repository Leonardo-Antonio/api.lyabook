package handler

import (
	"errors"
	"net/http"
	"strings"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/helper"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/valid"
	"github.com/Leonardo-Antonio/validmor"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type book struct {
	storage model.Ibook
}

func NewBook(storage model.Ibook) *book {
	return &book{storage}
}

func (b *book) Create(ctx echo.Context) error {
	valid := valid.Book()

	slug := ctx.Param("format")

	var book entity.Book
	if err := ctx.Bind(&book); err != nil {
		return response.New(ctx, http.StatusBadRequest, "la estructura no es valida", true, nil)
	}

	book.Slug = strings.ToLower(strings.Join(strings.Split(book.Name, " "), "-"))
	errs := validmor.ValidateStruct(book)
	if ers := valid.Format(&book.Type, slug); len(ers) != 0 {
		errs = append(errs, ers...)
	}
	if len(book.Categories) == 0 {
		errs = append(errs, errors.New("el libro debe tener al menos una categoria"))
	}
	if len(errs) != 0 {
		return response.New(ctx, http.StatusBadRequest, helper.ErrToString(errs), true, nil)
	}

	valid.CreateBook(&book)

	result, err := b.storage.Insert(&book)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return response.New(
				ctx, http.StatusBadRequest,
				"el nombre del libro <"+book.Name+"> ya existe",
				true, nil,
			)
		}
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	return response.New(ctx, http.StatusCreated, "el libro <"+book.Name+"> se creo correctamente", false, result)
}
