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

func (d *documentation) Index(ctx echo.Context) error {
	return response.New(ctx, http.StatusOK, "todo good", false, map[string][]string{
		"routers": {
			"/users/sign-in/:type",
			"/users/log-in/:type",
		},
	})
}
