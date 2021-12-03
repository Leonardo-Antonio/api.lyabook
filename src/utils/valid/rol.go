package valid

import (
	"errors"
	"strings"

	"github.com/Leonardo-Antonio/api.lyabook/src/utils/enum"
)

func Rol(rol string) error {
	if strings.ToUpper(rol) == enum.Rol.Manager ||
		strings.ToUpper(rol) == enum.Rol.Admin ||
		strings.ToUpper(rol) == enum.Rol.Client {
		return nil
	}

	return errors.New("el rol ingresado no es valido")
}
