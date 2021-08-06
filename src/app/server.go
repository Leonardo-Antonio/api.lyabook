package app

import (
	"log"

	"github.com/Leonardo-Antonio/api.lyabook/src/router"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/Leonardo-Antonio/api.lyabook/src/utils/key"
	"github.com/Leonardo-Antonio/gobcrypt/gobcrypt"
	"github.com/Leonardo-Antonio/validmor"
	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type server struct {
	app *echo.Echo
}

func New() *server {
	return &server{
		app: echo.New(),
	}
}
func (s *server) Middlewares() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.CORS())
}

func (s *server) Configs() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	env.GetEnv()

	key.Password = gobcrypt.CreateHash(env.Data.SecretKey)
	validmor.Errors(validmor.ERR_ES)
}

func (s *server) Routers() {
	router.Documentation(s.app)
}

func (s *server) Listeing() {
	s.app.Logger.Fatal(s.app.Start(env.Data.Port))
}
