package handler

import (
	"net/http"
	"strings"
	"time"

	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/tmpl"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/labstack/echo/v4"
)

type report struct{}

func NewReport() *report {
	return &report{}
}

func (r *report) Create(ctx echo.Context) error {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	tpl, err := tmpl.Report("stock.tpl", time.Now().Format(time.RFC1123))
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
