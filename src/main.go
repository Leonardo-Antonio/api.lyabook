package main

import (
	"github.com/Leonardo-Antonio/api.lyabook/src/app"
)

func main() {
	app := app.New(app.DEV)
	app.Configs()
	app.Middlewares()
	app.Routers()
	app.Listeing()
}
