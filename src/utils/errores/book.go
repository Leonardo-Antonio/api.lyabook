package errores

import "errors"

var (
	ErrFormatNotValid = errors.New("el formato que ingreso no es valido, ingrese <d> si es digital, <f> si es fisico, o <df> si quiere registrar ambos")
)
