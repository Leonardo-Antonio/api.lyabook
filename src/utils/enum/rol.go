package enum

type rol struct {
	Admin   string
	Manager string
	Client  string
}

var Rol = &rol{
	Admin:   "ADMIN",
	Manager: "MANAGER",
	Client:  "CLIENT",
}
