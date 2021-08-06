package main

import (
	"fmt"
	"log"

	"github.com/Leonardo-Antonio/api.lyabook/src/utils/env"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatalln(err)
	}
	env.GetEnv()
	fmt.Println(env.Data.BaseUrl)
}


