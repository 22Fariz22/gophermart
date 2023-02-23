package main

import (
	"github.com/22Fariz22/gophermart/internal/app"
	"github.com/22Fariz22/gophermart/internal/config"
	"log"
	"os"
)

func main() {
	cfg := config.NewConfig()

	log.Println("cfg from main: ", cfg)

	log.Println("len", len(os.Args))
	for _, arg := range os.Args[1:] {
		log.Println("arg: ", arg)
	}
	app := app.NewApp(cfg)

	app.Run()
}
