package handler

import (
	"net/http"

	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/labstack/echo/v4"
)

type documentation struct{}

func NewDocumentation() *documentation {
	return &documentation{}
}

func (d *documentation) Doc(ctx echo.Context) error {
	return ctx.Redirect(http.StatusPermanentRedirect, "https://documenter.getpostman.com/view/16524197/U16onNMr")
}

func (d *documentation) Index(ctx echo.Context) error {
	return response.New(ctx, http.StatusOK, "bienvenido a la api de lyabook develop", false, nil)
}
