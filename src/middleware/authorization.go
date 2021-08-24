package middleware

import (
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

		/* 		var i interface{}
		   		if err := ctx.Bind(&i); err != nil {
		   			log.Fatalln(err)
		   		}

		   		log.Println(i) */

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
