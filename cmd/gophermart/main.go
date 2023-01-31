package main

import (
	"github.com/22Fariz22/gophermart/internal/app"
	"github.com/22Fariz22/gophermart/internal/config"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("%s", err.Error())
	}

	app.Run(cfg)
}
