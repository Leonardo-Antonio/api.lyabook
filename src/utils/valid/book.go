package valid

import (
	"errors"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/enum"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/errores"
)

type book struct{}

func Book() *book {
	return &book{}
}

func (b *book) CreateBook(book *entity.Book) {
	book.Commentaries = nil
	book.PriceBefore = 0
	book.Stars = 0
}

func (b *book) Format(format *entity.Format, slug string) []error {
	switch slug {
	case "d":
		if errs := b.Digital(format); len(errs) != 0 {
			return errs
		}
	case "f":
		if errs := b.Fisico(format); len(errs) != 0 {
			return errs
		}
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
	if len(format.Digital.Details) == 0 {
		errs = append(errs, errors.New("el libro debe tener por lo menos un detalle"))
	}

	if len(errs) != 0 {
		return errs
	}

	return nil
}

func (b *book) Fisico(format *entity.Format) (errs []error) {
	format.Fisico.Format = enum.Format.Fisico
	if len(format.Fisico.Details) == 0 {
		errs = append(errs, errors.New("el libro debe tener por lo menos un detalle"))
	}
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
