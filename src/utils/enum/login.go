package enum

type login struct {
	Email string
	Dni   string
	Admin string
}

var TypeLogin = login{
	Email: "email",
	Dni:   "dni",
	Admin: "admin",
}
