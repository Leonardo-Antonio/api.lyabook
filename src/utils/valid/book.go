package valid

import (
	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
)

func CreateBook(book *entity.Book) {
	book.Commentaries = nil
	book.PriceBefore = 0
}
