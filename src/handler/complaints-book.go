package handler

import (
	"net/http"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/helper"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/Leonardo-Antonio/validmor"
	"github.com/labstack/echo/v4"
)

type complaintsBook struct {
	storage model.IComplaintsBook
}

func NewComplaintsBook(storage model.IComplaintsBook) *complaintsBook {
	return &complaintsBook{storage}
}

func (c *complaintsBook) Add(ctx echo.Context) error {
	var claim entity.ComplaintsBook
	if err := ctx.Bind(&claim); err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			err.Error(),
			true, nil,
		)
	}

	mapErrs := make(map[string][]string)
	errsClaimData := validmor.ValidateStruct(claim.ClaimData)
	errsDataPerson := validmor.ValidateStruct(claim.DataPerson)
	errsDataSendingReply := validmor.ValidateStruct(claim.DataSendingReply)

	if len(errsClaimData) != 0 {
		mapErrs["claim_data"] = helper.ErrToString(errsClaimData)
	}

	if len(errsDataPerson) != 0 {
		mapErrs["data_person"] = helper.ErrToString(errsDataPerson)
	}

	if len(errsDataSendingReply) != 0 {
		mapErrs["data_sending_reply"] = helper.ErrToString(errsDataSendingReply)
	}

	if len(mapErrs) != 0 {
		return response.New(
			ctx, http.StatusBadRequest,
			mapErrs,
			true, nil,
		)
	}

	result, err := c.storage.InsertOne(&claim)
	if err != nil {
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(),
			true, nil,
		)
	}

	return response.New(ctx, http.StatusCreated, "el reglamo se genero exitosamente", false, result)
}

func (c *complaintsBook) GetAll(ctx echo.Context) error {
	claims, err := c.storage.FindAll()
	if err != nil {
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(),
			true, nil,
		)
	}

	return response.New(ctx, http.StatusOK, "ok", false, claims)
}
