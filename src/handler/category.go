package handler

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/helper"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/formatting"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/Leonardo-Antonio/validmor"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type category struct {
	storage model.ICategory
}

func NewCategory(storage model.ICategory) *category {
	return &category{storage}
}

func (c *category) Add(ctx echo.Context) error {
	var category entity.Category
	if err := ctx.Bind(&category); err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			"la estructura que ingreso no es valida",
			true, nil,
		)
	}

	var err error
	if len(category.Slug) != 0 {
		err = errors.New("el campo slug se autogenera al crear la categoria")
		if len(category.Name) == 0 {
			return response.New(
				ctx, http.StatusBadRequest,
				helper.ErrToString([]error{err, errors.New("la categoria es un campo obligatorio y requerido")}),
				true, nil,
			)
		}
	}

	category.Slug = category.Name
	formatting.ReplaceSpecialCharacters(&category.Slug)

	errs := validmor.ValidateStruct(category)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) != 0 {
		return response.New(
			ctx, http.StatusBadRequest,
			helper.ErrToString(errs),
			true, nil,
		)
	}

	category.Name = formatting.ToTitle(category.Name)
	result, err := c.storage.Insert(&category)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return response.New(
				ctx, http.StatusBadRequest,
				"la categoria que quiere crear ya existe",
				true, nil,
			)
		}
		return response.New(
			ctx, http.StatusBadRequest,
			err.Error(),
			true, nil,
		)
	}

	return response.New(
		ctx, http.StatusCreated,
		fmt.Sprintf("la categoria <%s> se creo correctamente", category.Name),
		false, result,
	)
}

func (c *category) AddMany(ctx echo.Context) error {
	var categories []*entity.Category
	if err := ctx.Bind(&categories); err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			"la estructura que ingreso no es valida",
			true, nil,
		)
	}

	var errs []error
	for _, category := range categories {
		if len(category.Slug) != 0 {
			errs = append(errs, errors.New("el slug se genera automaticamente"))
		}

		category.Slug = category.Name
		formatting.ReplaceSpecialCharacters(&category.Slug)

		erros := validmor.ValidateStruct(*category)
		if len(erros) != 0 {
			errs = append(errs, erros...)
		}

		category.Name = formatting.ToTitle(category.Name)
	}

	if len(errs) != 0 {
		return response.New(
			ctx, http.StatusBadRequest,
			helper.ErrToString(errs),
			true, nil,
		)
	}

	result, err := c.storage.InsertMany(categories)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return response.New(
				ctx, http.StatusBadRequest,
				"verifique que las categorias no existan",
				true, nil,
			)
		}
		return response.New(
			ctx, http.StatusBadRequest,
			err.Error(),
			true, nil,
		)
	}

	return response.New(
		ctx, http.StatusCreated,
		"las categorias se crearon correctamente",
		false, result,
	)
}

func (c *category) Update(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			"el id ingresado no es valido",
			true, nil,
		)
	}

	var category entity.Category
	if err := ctx.Bind(&category); err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			"la estructura no es valida",
			true, nil,
		)
	}

	if len(category.Slug) != 0 {
		err = errors.New("el campo slug se autogenera al crear la categoria")
		if len(category.Name) == 0 {
			return response.New(
				ctx, http.StatusBadRequest,
				helper.ErrToString([]error{err, errors.New("la categoria es un campo obligatorio y requerido")}),
				true, nil,
			)
		}
	}

	category.Slug = category.Name
	formatting.ReplaceSpecialCharacters(&category.Slug)
	category.Id = id

	errs := validmor.ValidateStruct(category)
	if err != nil {
		errs = append(errs, err)
	}

	if len(errs) != 0 {
		return response.New(
			ctx, http.StatusBadRequest,
			helper.ErrToString(errs),
			true, nil,
		)
	}

	category.Name = formatting.ToTitle(category.Name)
	result, err := c.storage.Update(&category)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return response.New(
				ctx, http.StatusBadRequest,
				"la categoria que quiere crear ya existe",
				true, nil,
			)
		}
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(),
			true, nil,
		)
	}

	var message string
	if result.MatchedCount != 1 {
		message = "no hubo un cambio de los datos"
	} else {
		message = "se actualizo correctamente"
	}
	return response.New(
		ctx, http.StatusOK,
		message,
		false, result,
	)
}

func (c *category) DeleteById(ctx echo.Context) error {
	if len(ctx.QueryParam("id")) == 0 {
		return response.New(ctx, http.StatusBadRequest, "ingrese un id", true, nil)
	}

	id, err := primitive.ObjectIDFromHex(ctx.QueryParam("id"))
	if err != nil {
		return response.New(ctx, http.StatusBadRequest, fmt.Sprintf("el id <%s> no es valido", ctx.QueryParam("id")), true, nil)
	}

	result, err := c.storage.DeleteById(id)
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	if result.MatchedCount != 1 {
		return response.New(
			ctx, http.StatusBadRequest,
			fmt.Sprintf("no se encontro el libro con el id <%s>", ctx.QueryParam("id")),
			true, nil)
	}

	return response.New(
		ctx, http.StatusOK,
		"el libro se elimino correctamente",
		true, nil,
	)
}
