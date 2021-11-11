package handler

import (
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/response"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/tmpl"
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type payment struct {
	storage     model.IPayment
	storageBook model.Ibook
}

func NewPayment(storage model.IPayment, storageBook model.Ibook) *payment {
	return &payment{storage, storageBook}
}

func (p *payment) GetById(c echo.Context) error {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		return response.New(c, http.StatusBadRequest, err.Error(), true, nil)
	}
	data, err := p.storage.GetById(id)
	if err != nil {
		return response.New(c, http.StatusInternalServerError, err.Error(), true, nil)
	}

	total := 0
	for _, payment := range data {
		for _, item := range payment.Products {
			item.Importe = item.PriceUnit * float32(item.Quantity)
			total += int(item.Importe)
		}
	}

	data[0].TotalPagar = float32(total)
	data[0].CreateAtString = data[0].CreateAt.Format(time.RFC822)

	tpl, err := tmpl.Report("boleta.tpl", data)
	if err != nil {
		return response.New(c, http.StatusInternalServerError, err.Error(), true, nil)
	}
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		return response.New(c, http.StatusInternalServerError, err.Error(), true, nil)
	}
	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(tpl)))
	err = pdfg.Create()
	if err != nil {
		return response.New(c, http.StatusInternalServerError, err.Error(), true, nil)
	}

	err = pdfg.WriteFile("reports/boleta.pdf")
	if err != nil {
		return response.New(c, http.StatusInternalServerError, err.Error(), true, nil)
	}

	return c.File("reports/boleta.pdf")
}

func (p *payment) GetAllBooksSold(c echo.Context) error {
	dataBooksSold, err := p.storage.GetAllBooksSold()
	if err != nil {
		return response.New(c, http.StatusInternalServerError, err.Error(), true, nil)
	}

	var books []*entity.BookPayment
	for _, bookSold := range dataBooksSold {
		book, err := p.storageBook.FindByName(bookSold.Id)
		if err != nil {
			if errors.Is(err, mongo.ErrNoDocuments) {
				return response.New(c, http.StatusNoContent, err.Error(), true, nil)
			}
			return response.New(c, http.StatusInternalServerError, err.Error(), true, nil)
		}
		bookPayment := &entity.BookPayment{
			SoldBook: *bookSold,
			Data:     book,
		}
		books = append(books, bookPayment)
	}

	return response.New(c, http.StatusOK, "ok", false, books)
}
