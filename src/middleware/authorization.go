package middleware

import (
	"log"
	"net/http"
	"strings"

	"github.com/Leonardo-Antonio/api.lyabook/src/authorization"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/enum"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/labstack/echo/v4"
)

type auth struct{}

func Authorization() *auth {
	return &auth{}
}

func (a *auth) Admin(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("Authorization")
		claimUser, err := authorization.ValidateToken(token)
		if err != nil {
			return response.New(
				ctx, http.StatusForbidden,
				err.Error(),
				true, nil,
			)
		}

		log.Println("********************")
		log.Println(claimUser.Rol)
		log.Println(enum.Rol.Admin)
		log.Println("********************")

		if !strings.EqualFold(claimUser.Rol, enum.Rol.Admin) {
			return response.New(
				ctx, http.StatusUnauthorized,
				claimUser.Name+", usted no tiene acceso a esta sección",
				true, nil,
			)
		}
		return f(ctx)
	}
}

func (a *auth) Client(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("Authorization")
		claimUser, err := authorization.ValidateToken(token)
		if err != nil {
			return response.New(
				ctx, http.StatusForbidden,
				err.Error(),
				true, nil,
			)
		}

		if !strings.EqualFold(claimUser.Rol, enum.Rol.Client) {
			return response.New(
				ctx, http.StatusUnauthorized,
				claimUser.Name+", usted no tiene acceso a esta sección",
				true, nil,
			)
		}
		return f(ctx)
	}
}

func (a *auth) Manager(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("Authorization")
		claimUser, err := authorization.ValidateToken(token)
		if err != nil {
			return response.New(
				ctx, http.StatusForbidden,
				err.Error(),
				true, nil,
			)
		}

		log.Println("********************")
		log.Println(claimUser.Rol)
		log.Println(enum.Rol.Manager)
		log.Println("********************")

		if !strings.EqualFold(claimUser.Rol, enum.Rol.Manager) {
			return response.New(
				ctx, http.StatusUnauthorized,
				claimUser.Name+", usted no tiene acceso a esta sección",
				true, nil,
			)
		}
		return f(ctx)
	}
}

func (a *auth) ManagerAndAdmin(f echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token := ctx.Request().Header.Get("Authorization")
		claimUser, err := authorization.ValidateToken(token)
		if err != nil {
			return response.New(
				ctx, http.StatusForbidden,
				err.Error(),
				true, nil,
			)
		}

		if !strings.EqualFold(claimUser.Rol, enum.Rol.Client) {
			return f(ctx)

		}

		return response.New(
			ctx, http.StatusUnauthorized,
			claimUser.Name+", usted no tiene acceso a esta sección",
			true, nil,
		)
	}
}
