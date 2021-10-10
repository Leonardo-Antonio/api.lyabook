package handler

import (
	"fmt"
	"net/http"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/send"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/tmpl"
	"github.com/labstack/echo/v4"
)

type manager struct {
}

func NewManager() *manager {
	return &manager{}
}

func (m *manager) MessageAdmin(ctx echo.Context) error {
	var data entity.MessageEmail
	if err := ctx.Bind(&data); err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			err.Error(),
			true, nil,
		)
	}

	tpl, err := tmpl.Read("message-admin", data)
	if err != nil {
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(),
			true, nil,
		)
	}

	if err := send.SendMany(tpl, data.Subject, data.From); err != nil {
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(),
			true, nil,
		)
	}

	return response.New(ctx, http.StatusOK, fmt.Sprintf("mensaje enviado con exito a %s", data.Name), false, nil)
}
