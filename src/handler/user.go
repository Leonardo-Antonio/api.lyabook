package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/Leonardo-Antonio/api.lyabook/src/authorization"
	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/helper"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/enum"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/errores"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/formatting"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/key"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/reniec"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/send"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/tmpl"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/valid"
	"github.com/Leonardo-Antonio/gobcrypt/gobcrypt"
	"github.com/Leonardo-Antonio/validmor"
	"github.com/labstack/echo/v4"
	"github.com/twharmon/gouid"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type user struct {
	storage model.IUser
}

func NewUser(storage model.IUser) *user {
	return &user{storage}
}

func (u *user) SearchById(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return response.New(ctx, http.StatusBadRequest, "el id ingresado no es valido", true, nil)
	}
	data, err := u.storage.FindById(id)
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	return response.New(ctx, http.StatusOK, "ok", false, data)
}

func (u *user) SearchDni(ctx echo.Context) error {
	dni := ctx.Param("dni")
	if len(dni) != 8 {
		return response.New(ctx, http.StatusBadRequest, "el dni no es valido", true, nil)
	}

	user := &entity.User{Dni: dni}
	err := reniec.GetDniReniec(user)
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	user.Name = strings.Title(strings.ToLower(user.Name))
	user.LastName = strings.Title(strings.ToLower(user.LastName))

	return response.New(ctx, http.StatusOK, "dni encontrado con exito", false, user)
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
		userData.Active = true

	case enum.TypeLogin.Email:
		if err := valid.Email(&userData); err != nil {
			errs = append(errs, err)
		}

		if len(errs) != 0 {
			return response.New(ctx, http.StatusBadRequest, helper.ErrToString(errs), true, nil)
		}
		userData.Active = false
		userData.VerificationCode = gouid.String(8, gouid.UpperCaseAlphaNum)

	default:
		return response.New(
			ctx, http.StatusBadRequest,
			"el tipo de registro no es valido, ingrese email o dni",
			true, nil,
		)
	}

	formatting.User(&userData)
	userData.Id = primitive.NewObjectID()
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

	resp := "el usuario <" + userData.Dni + "> fue creado correctamente"
	if flag == enum.TypeLogin.Email {
		resp = "el usuario <" + userData.Email + "> fue creado correctamente"

		verify := &entity.VerificationAccountAdmin{
			Name: userData.Name,
			Link: fmt.Sprintf(
				"%sverificacion-cuenta?id=%s&code=%s",
				env.Data.AppClient,
				userData.Id.Hex(),
				userData.VerificationCode,
			),
		}

		tpl, err := tmpl.Read("new-admin", verify)
		if err != nil {
			return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
		}

		if err := send.SendMany(tpl, fmt.Sprintf("Hola %s, verifica tu cuenta", userData.Name), userData.Email); err != nil {
			return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
		}
	}

	if flag == enum.TypeLogin.Admin {
		resp = fmt.Sprintf("el administrador <%s>, se creo correctamente", userData.Name)

		verify := &entity.VerificationAccountAdmin{
			Name: userData.Name,
			Link: fmt.Sprintf("%s/verify/admin/%s/%s", env.Data.AppClient, userData.Id.Hex(), userData.VerificationCode),
		}

		tpl, err := tmpl.Read("new-admin", verify)
		if err != nil {
			return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
		}

		if err := send.SendMany(tpl, fmt.Sprintf("Hola %s, verifica tu cuenta", userData.Name), userData.Email); err != nil {
			return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
		}
	}

	return response.New(ctx, http.StatusCreated,
		resp,
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
					"el usuaio <"+credentials.Dni+"> no existe o esta inactivo",
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
					"el usuaio <"+credentials.Email+"> no existe o falta validar el correo, para ello debe ingresar a su correo y dar clic al boton 'validar' que se encuentra en el mensaje que se le envío.",
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

	token, err := authorization.GenerateToken(&dataUser)
	if err != nil {
		response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	resp := make(map[string]interface{}, 2)
	resp["token"] = token
	resp["user"] = dataUser

	return response.New(ctx, http.StatusOK, "ok", false, resp)
}

func (u *user) ValidPasswordDecript(passEncrypt, pass string) (bool, error) {
	passDecode, err := gobcrypt.Decrypt(passEncrypt, []byte(key.Password))
	if err != nil {
		return false, err
	}

	if passDecode != pass {
		return false, errores.ErrInvalidPassword
	}

	return true, nil
}

func (u *user) VerifyAccount(ctx echo.Context) error {
	var userData entity.User
	if err := ctx.Bind(&userData); err != nil {
		return response.New(ctx, http.StatusBadRequest, err.Error(), true, nil)
	}

	if err := valid.Email(&userData); err != nil {
		return response.New(ctx, http.StatusBadRequest, err.Error(), true, nil)
	}

	if len(userData.VerificationCode) != 8 {
		return response.New(ctx, http.StatusBadRequest, "el codigo no es valido", true, nil)
	}

	data, err := u.storage.VerifyAccount(
		userData.Email,
		userData.VerificationCode,
	)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return response.New(
				ctx, http.StatusBadRequest,
				"el código ingresado no es valido <"+userData.VerificationCode+">",
				true, nil,
			)
		}
		return response.New(ctx, http.StatusBadRequest, err.Error(), true, nil)
	}

	if err := send.EmailSignUp(data); err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	return response.New(
		ctx, http.StatusCreated,
		"se valido y creo correctamente su cuenta <"+userData.Email+">",
		false, nil,
	)
}

func (u *user) VerifyAccountById(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return response.New(ctx, http.StatusBadRequest, "el id no es valido", true, nil)
	}
	code := ctx.Param("code")

	data, err := u.storage.VerifyAccountById(id, code)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return response.New(
				ctx, http.StatusBadRequest,
				"el código ingresado no es valido <"+code+">",
				true, nil,
			)
		}
		return response.New(ctx, http.StatusBadRequest, err.Error(), true, nil)
	}

	if err := send.EmailSignUp(data); err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	return response.New(
		ctx, http.StatusCreated,
		"se valido y creo correctamente su cuenta",
		false, nil,
	)
}

func (u *user) Update(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			"el id ingresado no es valido",
			true, nil,
		)
	}

	var user entity.User
	if err := ctx.Bind(&user); err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			"la estructura que ingreso no es correcta",
			true, nil,
		)
	}

	user.Id = id
	user.Dni = ""
	user.Email = ""

	errs := validmor.ValidateStruct(user)
	if user.Id.IsZero() {
		errs = append(errs, errors.New("el id no debe ser nulo ni vacio, _id es un campo requerido"))
	}

	if len(errs) != 0 {
		return response.New(
			ctx, http.StatusBadRequest,
			helper.ErrToString(errs),
			true, nil,
		)
	}

	user.Password, err = gobcrypt.Encrypt(user.Password, []byte(key.Password))
	if err != nil {
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(),
			true, nil,
		)
	}

	u.validateFieldsNoUpdated(&user)
	result, err := u.storage.Update(&user)
	if err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			err.Error(),
			true, nil,
		)
	}

	if result.MatchedCount != 1 {
		return response.New(ctx, http.StatusBadRequest, "No se logro actualzar, verifique que todo este bien", true, nil)
	}

	return response.New(ctx, http.StatusOK, "se acualizo correctamente", false, result)
}

func (u *user) UpdateDataPerson(ctx echo.Context) error {
	id, err := primitive.ObjectIDFromHex(ctx.Param("id"))
	if err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			"el id ingresado no es valido",
			true, nil,
		)
	}

	var user entity.User
	if err := ctx.Bind(&user); err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			"la estructura que ingreso no es correcta",
			true, nil,
		)
	}

	user.Id = id
	user.Dni = ""
	user.Email = ""
	user.Password = ""

	errs := validmor.ValidateStruct(user)
	if user.Id.IsZero() {
		errs = append(errs, errors.New("el id no debe ser nulo ni vacio, _id es un campo requerido"))
	}

	if len(errs) != 0 {
		return response.New(
			ctx, http.StatusBadRequest,
			helper.ErrToString(errs),
			true, nil,
		)
	}

	u.validateFieldsNoUpdated(&user)
	result, err := u.storage.Update(&user)
	if err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			err.Error(),
			true, nil,
		)
	}

	if result.MatchedCount != 1 {
		return response.New(ctx, http.StatusBadRequest, "No se logro actualzar, verifique que todo este bien", true, nil)
	}

	return response.New(ctx, http.StatusOK, "se actualizaron su datos personales", false, result)
}

func (u *user) validateFieldsNoUpdated(user *entity.User) {
	user.Rol = ""
	user.Dni = ""
	user.VerificationCode = ""
}

func (u *user) FindAllUsersByRol(ctx echo.Context) error {
	admins, err := u.storage.FindAllUsersByRol(strings.Title(strings.ToLower(enum.Rol.Admin)))
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	return response.New(ctx, http.StatusOK, "ok", false, admins)
}

func (u *user) DeleteById(ctx echo.Context) error {
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

	result, err := u.storage.DeleteById(id)
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	if result.MatchedCount != 1 {
		return response.New(
			ctx, http.StatusBadRequest,
			fmt.Sprintf("no se encontro el admin con el id <%s>", ctx.QueryParam("id")),
			true, nil)
	}

	return response.New(
		ctx, http.StatusOK,
		"el admin se elimino correctamente",
		false, result,
	)
}

func (u *user) CountByRol(ctx echo.Context) error {
	amount, err := u.storage.CountUserByRol(strings.Title(strings.ToLower(enum.Rol.Client)))
	if err != nil {
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(), true, nil,
		)
	}

	return response.New(
		ctx, http.StatusOK,
		"ok", false, amount,
	)
}
