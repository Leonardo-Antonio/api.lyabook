package env

import (
	"os"
	"sync"
)

type (
	env struct {
		BaseUrl,
		UrlMongo,
		DBName,
		SecretKey,
		ApiReniecDni,
		TokenReniecDni,
		Port string
	}
)

var once sync.Once
var Data *env

func GetEnv() {
	once.Do(func() {
		Data = &env{
			BaseUrl:        os.Getenv("BASE_URI"),
			UrlMongo:       os.Getenv("URL_MONGO"),
			DBName:         os.Getenv("DB_NAME"),
			SecretKey:      os.Getenv("SECRET_KEY"),
			Port:           ":" + os.Getenv("PORT"),
			ApiReniecDni:   os.Getenv("API_RENIEC_DNI"),
			TokenReniecDni: os.Getenv("TOKEN_API_RENIEC_DNI"),
		}
	})
}
