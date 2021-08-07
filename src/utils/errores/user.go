package errores

import "errors"

var (
	ErrTypeLogin       = errors.New("el tipo de registro o ingreso no es valido, pruebe con <email, dni>")
	ErrInvalidPassword = errors.New("el password ingresado no es valido")
)
