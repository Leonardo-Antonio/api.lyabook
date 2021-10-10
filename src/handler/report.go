package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/tmpl"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/mongo"
)

type report struct {
	storage model.Ibook
}

func NewReport(storage model.Ibook) *report {
	return &report{storage}
}

func (r *report) AllBooks(ctx echo.Context) error {
	books, err := r.storage.FindByFormat("df")
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	tpl, err := tmpl.Report("stock.tpl", books) // time.Now().Format(time.RFC1123)
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(tpl)))

	err = pdfg.Create()
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	err = pdfg.WriteFile("reports/all-books_stock.pdf")
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	currentData := time.Now()
	return ctx.Attachment(
		"reports/stock.pdf",
		fmt.Sprintf("%d-%d-%d.pdf", currentData.Day(), currentData.Month(), currentData.Year()),
	)
}

func (r *report) SearchByFormat(ctx echo.Context) error {
	format := ctx.Param("format")

	books, err := r.storage.FindByFormat(format)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return response.New(
				ctx, http.StatusNoContent,
				"no hay datos",
				true, nil,
			)
		}
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(),
			true, nil,
		)
	}

	return response.New(
		ctx, http.StatusOK,
		"ok",
		false, books,
	)
}
