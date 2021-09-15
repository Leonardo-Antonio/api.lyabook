package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type documentation struct{}

func NewDocumentation() *documentation {
	return &documentation{}
}

func (d *documentation) Index(ctx echo.Context) error {
	return ctx.Redirect(http.StatusPermanentRedirect, "https://documenter.getpostman.com/view/16524197/U16onNMr")
}
