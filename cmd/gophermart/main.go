package main

import (
	"github.com/22Fariz22/gophermart/internal/app"
	"github.com/22Fariz22/gophermart/internal/config"
)

func main() {
	cfg := config.NewConfig()

	app := app.NewApp(cfg)

	app.Run()
}
