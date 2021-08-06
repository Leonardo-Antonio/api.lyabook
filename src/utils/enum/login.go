package enum

type login struct {
	Email string
	Dni   string
}

var TypeLogin = login{
	Email: "email",
	Dni:   "dni",
}
