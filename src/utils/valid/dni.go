package valid

import "errors"

func Dni(dni string) error {
	if len(dni) == 8 {
		return nil
	}
	return errors.New("el dni ingresado no es valido")
}
