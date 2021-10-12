package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
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
	var data entity.ReportBooks
	var err error
	data.Books, err = r.storage.FindByFormat("df")
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	data.Date = time.Now().Format(time.RFC1123)

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	tpl, err := tmpl.Report("all-books.tpl", data) // time.Now().Format(time.RFC1123)
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
		"reports/all-books_stock.pdf",
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

func (r *report) SearchBookByStock(ctx echo.Context) error {
	stock, err := strconv.Atoi(ctx.Param("stock"))
	if err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			fmt.Sprintf("el stock ingresado no es valido <%s>", ctx.Param("stock")),
			true, nil,
		)
	}

	books, err := r.storage.FindByStock(uint(stock))
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

func (r *report) ReportBooksByStock(ctx echo.Context) error {
	stock, err := strconv.Atoi(ctx.Param("stock"))
	if err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			fmt.Sprintf("el stock ingresado no es valido <%s>", ctx.Param("stock")),
			true, nil,
		)
	}

	var data entity.ReportBooks
	data.Books, err = r.storage.FindByStock(uint(stock))
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}
	data.Date = time.Now().Format(time.RFC1123)

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	tpl, err := tmpl.Report("stock.tpl", data) // time.Now().Format(time.RFC1123)
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(tpl)))

	err = pdfg.Create()
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	err = pdfg.WriteFile("reports/stock.pdf")
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	return ctx.File(
		"reports/stock.pdf",
	)
}

func (r *report) NewBooks(ctx echo.Context) error {
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			fmt.Sprintf("el limite que agrego no es valido <%s>", ctx.Param("limit")),
			true, nil,
		)
	}

	books, err := r.storage.NewBooksOfTheMonth(uint(limit))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return response.New(
				ctx, http.StatusNoContent,
				"no hay libros registrados",
				true, nil,
			)
		}
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(),
			true, nil,
		)
	}

	return response.New(ctx, http.StatusOK, "ok", false, books)
}

func (r *report) ReportNewBooks(ctx echo.Context) error {
	limit, err := strconv.Atoi(ctx.Param("limit"))
	if err != nil {
		return response.New(
			ctx, http.StatusBadRequest,
			fmt.Sprintf("el limite que agrego no es valido <%s>", ctx.Param("limit")),
			true, nil,
		)
	}

	var data entity.ReportBooks
	data.Books, err = r.storage.NewBooksOfTheMonth(uint(limit))
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return response.New(
				ctx, http.StatusNoContent,
				"no hay libros registrados",
				true, nil,
			)
		}
		return response.New(
			ctx, http.StatusInternalServerError,
			err.Error(),
			true, nil,
		)
	}
	data.Date = time.Now().Format(time.RFC1123)

	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	tpl, err := tmpl.Report("new-books.tpl", data) // time.Now().Format(time.RFC1123)
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(tpl)))

	err = pdfg.Create()
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	err = pdfg.WriteFile("reports/new-books.pdf")
	if err != nil {
		return response.New(ctx, http.StatusInternalServerError, err.Error(), true, nil)
	}

	return ctx.File(
		"reports/new-books.pdf",
	)
}
