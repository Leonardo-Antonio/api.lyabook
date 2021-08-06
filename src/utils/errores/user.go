package errores

import "errors"

var (
	ErrTypeLogin = errors.New("the flag does not exist, enter one of the options: email or dni")
)
