package valid

import (
	"errors"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/enum"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/errores"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/formatting"
)

type book struct{}

func Book() *book {
	return &book{}
}

func (b *book) CreateBook(book *entity.Book) {
	book.Name = formatting.ToTitle(book.Name)
	book.Editorial = formatting.ToTitle(book.Editorial)
	book.Author = formatting.ToTitle(book.Author)
	book.Commentaries = nil
	book.PriceBefore = book.PriceCurrent
	book.Stars = 0
}

func (b *book) ValidArrays(book entity.Book) (errs []error) {
	if len(book.ImagesSrc) == 0 {
		errs = append(errs, errors.New("el libro debe tener al menos una imagen"))
	}

	if len(book.Details) == 0 {
		errs = append(errs, errors.New("el libro debe tener al menos detalle sobre el libro"))
	}

	if len(book.Categories) == 0 {
		errs = append(errs, errors.New("el libro debe tener al menos una categoria"))
	}

	return errs
}

func (b *book) Format(format *entity.Format, slug string) []error {
	switch slug {
	case "d":
		if errs := b.Digital(format); len(errs) != 0 {
			return errs
		}

		format.Fisico = entity.Fisico{}
	case "f":
		if errs := b.Fisico(format); len(errs) != 0 {
			return errs
		}

		format.Digital = entity.Digital{}
	case "df":
		if errs := b.Digital(format); len(errs) != 0 {
			return errs
		}
		if errs := b.Fisico(format); len(errs) != 0 {
			return errs
		}
	default:
		return []error{errores.ErrFormatNotValid}
	}

	return nil
}

func (b *book) Digital(format *entity.Format) (errs []error) {
	format.Digital.Format = enum.Format.Digital
	if len(format.Digital.Src) == 0 {
		errs = append(errs, errors.New("ingrese una url del pdf u ebook"))
	}

	if len(errs) != 0 {
		return errs
	}

	return nil
}

func (b *book) Fisico(format *entity.Format) (errs []error) {
	format.Fisico.Format = enum.Format.Fisico

	if len(format.Fisico.Lat) == 0 {
		errs = append(errs, errors.New("la lantitud <lat> no es valida"))
	}
	if len(format.Fisico.Log) == 0 {
		errs = append(errs, errors.New("la longitud <log> no es valida"))
	}
	if format.Fisico.Stock < 1 {
		errs = append(errs, errors.New("el stock no es valido"))
	}

	if len(errs) != 0 {
		return errs
	}

	return nil
}
