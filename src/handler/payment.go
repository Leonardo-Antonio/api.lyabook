package handler

import (
	"net/http"

	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type payment struct {
	storage model.IPayment
}

func NewPayment(storage model.IPayment) *payment {
	return &payment{storage}
}

func (p *payment) GetById(c echo.Context) error {
	id, _ := primitive.ObjectIDFromHex("6167d3817d2fd42d6f5fdd7b")
	data, _ := p.storage.GetById(id)
	return response.New(c, http.StatusOK, "ok", false, data)
}
