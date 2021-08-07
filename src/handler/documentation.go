package handler

import (
	"net/http"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type documentation struct{}

func NewDocumentation() *documentation {
	return &documentation{}
}

func (d *documentation) Index(ctx echo.Context) error {
	u := entity.User{
		Id:   primitive.NewObjectID(),
		Name: "Ñepd",
	}
	return response.New(ctx, http.StatusOK, "todo good", false, u)
}
