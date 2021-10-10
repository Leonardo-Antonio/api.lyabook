package app

import (
	"log"

	"github.com/Leonardo-Antonio/api.lyabook/src/authorization"
	"github.com/Leonardo-Antonio/api.lyabook/src/dbutil"
	"github.com/Leonardo-Antonio/api.lyabook/src/model"
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
	app   *echo.Echo
	state string
}

func New(state string) *server {
	return &server{
		app:   echo.New(),
		state: state,
	}
}
func (s *server) Middlewares() {
	s.app.Use(middleware.Recover())
	s.app.Use(middleware.Logger())
	s.app.Use(middleware.CORS())
}

func (s *server) Configs() {
	s.loadCertificates()
	if s.state == DEV {
		if err := godotenv.Load(); err != nil {
			log.Fatalln(err)
		}
	}
	env.GetEnv()

	key.Password = gobcrypt.CreateHash(env.Data.SecretKey)
	validmor.Errors(validmor.ERR_ES)
}

func (s *server) loadCertificates() {
	authorization.LoadFiles()
}

func (s *server) Routers() {
	db := dbutil.GetConnection()
	router.Documentation(s.app)
	router.User(model.NewUser(db), s.app)
	router.Category(model.NewCategoty(db), s.app)
	router.Book(model.NewBook(db), model.NewUser(db), model.NewCategoty(db), s.app)
	router.ComplaintsBook(model.NewComplaintsBook(db), s.app)
	router.Report(model.NewBook(db), s.app)
	router.Manager(s.app)
}

func (s *server) Listeing() {
	s.app.Logger.Fatal(s.app.Start(env.Data.Port))
}
