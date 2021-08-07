package handler

import (
	"errors"
	"net/http"

	"github.com/Leonardo-Antonio/api.lyabook/src/autorization"
	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/helper"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/enum"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/errores"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/formatting"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/key"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/valid"
	"github.com/Leonardo-Antonio/gobcrypt/gobcrypt"
	"github.com/Leonardo-Antonio/validmor"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	storage model.IUser
}

func NewUser(storage model.IUser) *user {
	return &user{storage}
}

func (u *user) SignUp(ctx echo.Context) error {
	flag := ctx.Param("type")

	var userData entity.User
	if err := ctx.Bind(&userData); err != nil {
		return response.New(ctx, http.StatusBadRequest, err.Error(), true, nil)
	}

	errs := validmor.ValidateStruct(userData)
	passEncrypt, err := gobcrypt.Encrypt(userData.Password, []byte(key.Password))
	userData.Password = passEncrypt
	if err != nil {
		errs = append(errs, err)
	}

	if err := valid.Rol(userData.Rol); err != nil {
		errs = append(errs, err)
	}

	switch flag {
	case enum.TypeLogin.Dni:
		if err := valid.Dni(userData.Dni); err != nil {
			errs = append(errs, err)
		}

		if len(errs) != 0 {
			return response.New(ctx, http.StatusBadRequest, helper.ErrToString(errs), true, nil)
		}

		if err := helper.GetDniReniec(&userData); err != nil {
			if errors.Is(err, errores.ErrFindDniApiReniec) {
				return response.New(ctx, http.StatusBadRequest, err.Error(), true, nil)
			}
			return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
		}

	case enum.TypeLogin.Email:
		if err := valid.Email(&userData); err != nil {
			errs = append(errs, err)
		}

		if len(errs) != 0 {
			return response.New(ctx, http.StatusBadRequest, helper.ErrToString(errs), true, nil)
		}
	default:
		return response.New(
			ctx, http.StatusBadRequest,
			"el tipo de registro no es valido, ingrese email o dni",
			true, nil,
		)
	}

	formatting.User(&userData)
	result, err := u.storage.Insert(&userData)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			message := "el email <" + userData.Email + "> ya esta en uso, ingrese otro"
			if flag == enum.TypeLogin.Dni {
				message = "el dni <" + userData.Dni + "> ya esta en uso, ingrese otro"
			}
			return response.New(ctx, http.StatusBadRequest,
				message,
				true, nil,
			)
		}
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}
	return response.New(ctx, http.StatusCreated,
		"el usuario <"+userData.Dni+"> fue creado correctamente",
		false, result,
	)
}

func (u *user) LogIn(ctx echo.Context) error {
	flag := ctx.Param("type")

	var credentials entity.User
	if err := ctx.Bind(&credentials); err != nil {
		return response.New(ctx, http.StatusBadRequest, "la estructura ingresada no es valida", true, nil)
	}

	errs := validmor.ValidateStruct(credentials)

	var err error
	var dataUser entity.User
	switch flag {
	case enum.TypeLogin.Dni:
		if err := valid.Dni(credentials.Dni); err != nil {
			errs = append(errs, err)
		}

		if len(errs) != 0 {
			return response.New(ctx, http.StatusBadRequest, helper.ErrToString(errs), true, nil)
		}

		dataUser, err = u.storage.Find(&credentials, enum.TypeLogin.Dni)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return response.New(
					ctx, http.StatusBadRequest,
					"el usuaio <"+credentials.Dni+"> no existe o no ingreso correctamente su credenciales",
					true, nil,
				)
			}
			return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
		}

	case enum.TypeLogin.Email:
		if err := valid.Email(&credentials); err != nil {
			errs = append(errs, err)
		}

		if len(errs) != 0 {
			return response.New(ctx, http.StatusBadRequest, helper.ErrToString(errs), true, nil)
		}

		dataUser, err = u.storage.Find(&credentials, enum.TypeLogin.Email)
		if err != nil {
			if errors.Is(err, errores.ErrInvalidPassword) {
				return response.New(
					ctx, http.StatusBadRequest,
					"el password <"+credentials.Password+"> no es correcto o el usuario no existe",
					true, nil,
				)
			}
			if errors.Is(err, mongo.ErrNoDocuments) {
				return response.New(
					ctx, http.StatusBadRequest,
					"el usuaio <"+credentials.Email+"> no existe o no ingreso correctamente su credenciales",
					true, nil,
				)
			}
			return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
		}

	default:
		return response.
			New(
				ctx, http.StatusBadRequest,
				"el tipo de login no es valido, pruebe con <email, dni>",
				true, nil,
			)
	}

	token, err := autorization.GenerateToken(&dataUser)
	if err != nil {
		response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	resp := make(map[string]interface{}, 2)
	resp["token"] = token
	resp["user"] = dataUser

	return response.New(ctx, http.StatusOK, "ok", false, resp)
}
