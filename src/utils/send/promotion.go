package send

import (
	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
)

func Promotion(book entity.Book, to ...string) error {
	tpl, err := readTemplate("promotion.tpl", "template/promotion.tpl", book)
	if err != nil {
		return err
	}
	if err := sendMany(tpl, to...); err != nil {
		return err
	}
	return nil
}
