package env

import (
	"os"
	"sync"
)

type (
	env struct {
		BaseUrl, UrlMongo, DBName, SecretKey, Port string
	}
)

var once sync.Once
var Data *env

func GetEnv() {
	once.Do(func() {
		Data = &env{
			BaseUrl:   os.Getenv("BASE_URI"),
			UrlMongo:  os.Getenv("URL_MONGO"),
			DBName:    os.Getenv("DB_NAME"),
			SecretKey: os.Getenv("SECRET_KEY"),
			Port:      ":" + os.Getenv("PORT"),
		}
	})
}
