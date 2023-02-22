package main

import (
	"github.com/22Fariz22/gophermart/internal/app"
	"github.com/22Fariz22/gophermart/internal/config"
	"log"
)

func main() {
	_, err := config.NewConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	app := app.NewApp()

	app.Run()
}
