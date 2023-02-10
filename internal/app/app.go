package app

import (
	"context"
	"fmt"
	"github.com/22Fariz22/gophermart/config"
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/22Fariz22/gophermart/internal/auth/delivery"
	postgres2 "github.com/22Fariz22/gophermart/internal/auth/repository/postgres"
	"github.com/22Fariz22/gophermart/internal/auth/usecase"
	"github.com/22Fariz22/gophermart/internal/order"
	delivery2 "github.com/22Fariz22/gophermart/internal/order/delivery"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server

	authUC  auth.UseCase
	orderUC order.UseCase
}

func NewApp(cfg *config.Config) *App {

	// Repository
	db, err := postgres.New(cfg.DATABASE_URI)
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	defer db.Close()

	userRepo := postgres2.NewUserRepository(db)

	return &App{
		authUC: usecase.NewAuthUseCase(
			userRepo,
			"hash_salt",
			[]byte("signing_key"),
			time.Duration(86400),
		),
	}
}
func (a *App) Run() error {
	l := logger.New("debug")

	// Init gin handler
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// Set up http handlers
	// SignUp/SignIn endpoints
	delivery.RegisterHTTPEndpoints(router, a.authUC)

	// API endpoints
	authMiddleware := delivery.NewAuthMiddleware(a.authUC)
	api := router.Group("/api/user", authMiddleware)

	delivery2.RegisterHTTPEndpointsOrder(api, a.orderUC)

	// HTTP Server
	a.httpServer = &http.Server{
		Addr:           ":" + "8080",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := a.httpServer.ListenAndServe(); err != nil {
			l.Fatal("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return a.httpServer.Shutdown(ctx)
}
