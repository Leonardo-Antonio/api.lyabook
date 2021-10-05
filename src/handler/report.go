package handler

import (
	"net/http"
	"strings"

	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/tmpl"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/labstack/echo/v4"
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

	//Your Pdf Name
	err = pdfg.WriteFile("reports/stock.pdf")
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	return ctx.Attachment("reports/stock.pdf", "stock.pdf")
}
