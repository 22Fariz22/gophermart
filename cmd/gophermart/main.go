package main

import (
	"fmt"
	"github.com/22Fariz22/gophermart/internal/app"
	"github.com/22Fariz22/gophermart/internal/config"
	"os"
)

func main() {
	fmt.Println("len", len(os.Args))
	for _, arg := range os.Args[1:] {
		fmt.Println("arg: ", arg)
	}

	cfg := config.NewConfig()

	app := app.NewApp(cfg)

	app.Run()
}
