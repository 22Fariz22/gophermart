package app

import (
	"fmt"
	"github.com/22Fariz22/gophermart/config"
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"net/http"
)

type App struct {
	httpServer *http.Server

	userUC auth.UseCase
}

func Run(cfg *config.Config) {
	l := logger.New(cfg.LogLevel)

	// Repository
	pg, err := postgres.New(cfg)
	if err != nil {
		l.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer pg.Close()

	//gin.New()
}
