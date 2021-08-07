package valid

import (
	"errors"
	"regexp"

	"github.com/Leonardo-Antonio/api.lyabook/src/entity"
)

func Email(user *entity.User) error {
	if len(user.Email) == 0 {
		return errors.New("el email ingresado no es valido")
	}
	if !regexp.
		MustCompile(
			"^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$",
		).MatchString(user.Email) {
		return errors.New("el email ingresado no es valido")
	}

	return nil
}
