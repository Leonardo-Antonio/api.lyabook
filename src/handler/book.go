package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/helper"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/formatting"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/send"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/valid"
	"github.com/Leonardo-Antonio/validmor"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type book struct {
	storageBook     model.Ibook
	storageUser     model.IUser
	storageCategory model.ICategory
}

func NewBook(storageBook model.Ibook, storageUser model.IUser, storageCategory model.ICategory) *book {
	return &book{storageBook, storageUser, storageCategory}
}

func (b *book) Create(ctx echo.Context) error {
	valid := valid.Book()

	slug := ctx.Param("format")

	var book entity.Book
	if err := ctx.Bind(&book); err != nil {
		return response.New(ctx, http.StatusBadRequest, err.Error(), true, nil)
	}

	var idsInvalid []string
	for _, categotyEan := range book.Categories {
		_, err := b.storageCategory.SearchByEan(categotyEan)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				idsInvalid = append(idsInvalid, categotyEan)
			} else {
				return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
			}
		}
	}

	if len(idsInvalid) != 0 {
		return response.New(ctx, http.StatusBadRequest, "el 'ean' de las categorias ingresados no son validos o no existen", true, idsInvalid)
	}

	book.Slug = book.Name
	formatting.ReplaceSpecialCharacters(&book.Slug)

	errs := validmor.ValidateStruct(book)

	if len(book.ImagesSrc) == 0 {
		errs = append(errs, errors.New("el libro debe tener al menos una imagen"))
	}

	if len(book.Details) == 0 {
		errs = append(errs, errors.New("el libro debe tener al menos detalle sobre el libro"))
	}

	if ers := valid.Format(&book.Type, slug); len(ers) != 0 {
		errs = append(errs, ers...)
	}
	if len(book.Categories) == 0 {
		errs = append(errs, errors.New("el libro debe tener al menos una categoria"))
	}

	if len(errs) != 0 {
		return response.New(ctx, http.StatusBadRequest, helper.ErrToString(errs), true, nil)
	}

	book.FormatBook = slug
	valid.CreateBook(&book)
	result, err := b.storageBook.Insert(&book)
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

func (b *book) Edit(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			"el id <"+ctx.Param("id")+"> no es valido",
			true, nil)
	}

	valid := valid.Book()

	slug := ctx.Param("format")

	var book entity.Book
	if err := ctx.Bind(&book); err != nil {
		return response.New(ctx, http.StatusBadRequest, err.Error(), true, nil)
	}

	var idsInvalid []string
	for _, categotyEan := range book.Categories {
		_, err := b.storageCategory.SearchByEan(categotyEan)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				idsInvalid = append(idsInvalid, categotyEan)
			} else {
				return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
			}
		}
	}

	if len(idsInvalid) != 0 {
		return response.New(ctx, http.StatusBadRequest, "el 'ean' de las categorias ingresados no son validos o no existen", true, idsInvalid)
	}

	book.Slug = book.Name
	formatting.ReplaceSpecialCharacters(&book.Slug)

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

	book.Id = id
	result, err := b.storageBook.Update(&book)
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

	if result.MatchedCount != 1 {
		return response.New(
			ctx, http.StatusBadRequest,
			"no se logro actualizar el libro, el id no existe",
			true, nil,
		)
	}

	return response.New(ctx, http.StatusOK, "el libro <"+book.Name+"> se actualizo correctamente", false, result)
}

func (b *book) AddPromotion(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			"el id <"+id.String()+"> no es valido",
			true, nil,
		)
	}
	var book entity.Book
	if err := ctx.Bind(&book); err != nil {
		return response.New(ctx, http.StatusBadRequest, "la estructura no es valida", true, nil)
	}

	if book.PriceCurrent < 0 {
		return response.New(
			ctx, http.StatusBadRequest,
			"el precio no debe ser menos a 0",
			true, nil,
		)
	}

	dataBookPromotion, err := b.storageBook.FindBookById(id)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return response.New(
				ctx, http.StatusBadRequest,
				"el libro con el id <"+id.String()+"> no existe",
				true, nil,
			)
		}
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(),
			true, nil,
		)
	}

	if dataBookPromotion.PriceBefore < book.PriceCurrent {
		return response.New(
			ctx, http.StatusBadRequest,
			"no seria una oferta si agrega <"+strconv.Itoa(int(book.PriceCurrent))+"> ya que es mayor al precio anterior <"+strconv.Itoa(int(dataBookPromotion.PriceBefore))+">",
			true, nil,
		)
	}

	result, err := b.storageBook.UpdatePriceCurrent(id, book.PriceCurrent)
	if err != nil {
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(),
			true, nil,
		)
	}

	users, err := b.storageUser.FindUsersWithEmail()
	if err != nil {
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(),
			true, nil,
		)
	}

	if len(users) == 0 {
		return response.New(
			ctx, http.StatusOK,
			"se agrego la nueva promocion del libro",
			false, result,
		)
	}

	var emails []string
	for _, email := range users {
		emails = append(emails, email.Email)
	}

	log.Println(emails)

	if err := send.Promotion(dataBookPromotion, emails...); err != nil {
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(),
			true, nil,
		)
	}

	return response.New(
		ctx, http.StatusOK,
		"se agrego la nueva promocion del libro",
		false, result,
	)
}

func (b *book) DeleteById(ctx echo.Context) error {
	if len(ctx.QueryParam("id")) == 0 {
		return response.New(ctx, http.StatusBadRequest, "ingrese un id", true, nil)
	}

	id, err := primitive.ObjectIDFromHex(ctx.QueryParam("id"))
	if err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			fmt.Sprintf("el id <%s>no es valido",
				ctx.QueryParam("id")), true, nil,
		)
	}

	result, err := b.storageBook.DeleteById(id)
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
		false, result,
	)
}

func (b *book) CreateMany(ctx echo.Context) error {
	var books []*entity.Book
	if err := ctx.Bind(&books); err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			err.Error(),
			true, nil,
		)
	}

	for _, book := range books {

		var idsInvalid []string
		for _, categotyEan := range book.Categories {
			_, err := b.storageCategory.SearchByEan(categotyEan)
			if err != nil {
				if errors.Is(err, mongo.ErrNoDocuments) {
					idsInvalid = append(idsInvalid, categotyEan)
				} else {
					return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
				}
			}
		}

		if len(idsInvalid) != 0 {
			return response.New(ctx, http.StatusBadRequest, "el 'ean' de las categorias ingresados no son validos o no existen", true, idsInvalid)
		}

		var errs []error
		valid := valid.Book()
		valid.CreateBook(book)

		errs = validmor.ValidateStruct(*book)

		if errArrys := valid.ValidArrays(*book); len(errArrys) != 0 {
			errs = append(errs, errArrys...)
		}

		if (book.Type) == (entity.Format{}) {
			errs = append(errs, errors.New("el campo <type>, es obligatorio"))
		}

		if len(errs) != 0 {
			return response.New(
				ctx, http.StatusBadRequest,
				helper.ErrToString(errs), true, nil,
			)
		}

		errs = valid.Format(&book.Type, "df")
		if len(errs) != 0 {
			return response.New(
				ctx, http.StatusBadRequest,
				helper.ErrToString(errs), true, nil,
			)
		}

		book.Slug = book.Name
		formatting.ReplaceSpecialCharacters(&book.Slug)
		valid.CreateBook(book)
		book.FormatBook = "df"
	}

	result, err := b.storageBook.InsertMany(books)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return response.New(
				ctx, http.StatusBadRequest,
				err.Error(), true, nil,
			)
		}
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	return response.New(
		ctx, http.StatusCreated, "se crearon correctamente los libros", false, result,
	)
}
