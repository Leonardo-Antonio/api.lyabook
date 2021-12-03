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
	storage        model.Ibook
	storagePayment model.IPayment
	storageBook    model.Ibook
}

func NewReport(storage model.Ibook, storagePayment model.IPayment, storageBook model.Ibook) *report {
	return &report{storage, storagePayment, storageBook}
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

func (r *report) BooksSold(c echo.Context) error {
	dataBooksSold, err := r.storagePayment.GetAllBooksSold(0)
	if err != nil {
		return response.New(c, http.StatusInternalServerError, err.Error(), true, nil)
	}

	var max uint
	for _, item := range dataBooksSold {
		if item.BooksSold > max {
			max = item.BooksSold
		}
	}

	var books []*entity.BookPayment
	for _, bookSold := range dataBooksSold {
		book, err := r.storageBook.FindByName(bookSold.Id)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				continue
			}
			return response.New(c, http.StatusInternalServerError, err.Error(), true, nil)
		}

		var tag string
		if bookSold.BooksSold >= max/2 {
			tag = "Más vendido"
		} else {
			tag = "Menos vendido"
		}

		bookPayment := &entity.BookPayment{
			SoldBook:  *bookSold,
			Data:      book,
			Tag:       tag,
			CreatedAt: time.Now().Format(time.RFC1123),
		}
		books = append(books, bookPayment)
	}

	if err := r.GeneratePdf(books, "books-sold.tpl", "books-sold"); err != nil {
		return response.New(c, http.StatusInternalServerError, err.Error(), true, nil)
	}

	return c.File(
		"reports/books-sold.pdf",
	)
}

func (r *report) CategoriesSold(c echo.Context) error {
	dataBooksSold, err := r.storagePayment.GetAllBooksSold(0)
	if err != nil {
		return response.New(c, http.StatusInternalServerError, err.Error(), true, nil)
	}

	var max uint
	for _, item := range dataBooksSold {
		if item.BooksSold > max {
			max = item.BooksSold
		}
	}

	var books []*entity.BookPayment
	for _, bookSold := range dataBooksSold {
		book, err := r.storageBook.FindByName(bookSold.Id)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				continue
			}
			return response.New(c, http.StatusInternalServerError, err.Error(), true, nil)
		}

		var tag string
		if bookSold.BooksSold >= max/2 {
			tag = "Más vendido"
		} else {
			tag = "Menos vendido"
		}

		bookPayment := &entity.BookPayment{
			SoldBook:  *bookSold,
			Data:      book,
			Tag:       tag,
			CreatedAt: time.Now().Format(time.RFC1123),
		}
		books = append(books, bookPayment)
	}

	if err := r.GeneratePdf(books, "categories-sold.tpl", "categories-sold"); err != nil {
		return response.New(c, http.StatusInternalServerError, err.Error(), true, nil)
	}

	return c.File(
		"reports/categories-sold.pdf",
	)
}

func (r *report) GeneratePdf(data interface{}, nameTpl, namePdf string) error {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return err
	}

	tpl, err := tmpl.Report(nameTpl, data) // time.Now().Format(time.RFC1123)
	if err != nil {
		return err
	}
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(tpl)))

	err = pdfg.Create()
	if err != nil {
		return err
	}

	err = pdfg.WriteFile(fmt.Sprintf("reports/%s.pdf", namePdf))
	if err != nil {
		return err
	}
	return nil
}
