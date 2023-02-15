package app

import (
	"context"
	"fmt"
	"github.com/22Fariz22/gophermart/config"
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/22Fariz22/gophermart/internal/auth/delivery"
	postgres2 "github.com/22Fariz22/gophermart/internal/auth/repository/postgres"
	"github.com/22Fariz22/gophermart/internal/auth/usecase"
	"github.com/22Fariz22/gophermart/internal/history"
	delivery3 "github.com/22Fariz22/gophermart/internal/history/delivery"
	postgres4 "github.com/22Fariz22/gophermart/internal/history/repository/postgres"
	usecase3 "github.com/22Fariz22/gophermart/internal/history/usecase"
	"github.com/22Fariz22/gophermart/internal/order"
	delivery2 "github.com/22Fariz22/gophermart/internal/order/delivery"
	postgres3 "github.com/22Fariz22/gophermart/internal/order/repository/postgres"
	usecase2 "github.com/22Fariz22/gophermart/internal/order/usecase"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	httpServer *http.Server

	authUC    auth.UseCase
	orderUC   order.UseCase
	historyUC history.UseCase
}

func NewApp(cfg *config.Config) *App {

	// Repository
	db, err := postgres.New(viper.GetString("d"), postgres.MaxPoolSize(2))
	if err != nil {
		log.Fatal(fmt.Errorf("app - Run - postgres.New: %w", err))
	}
	//defer db.Close()

	userRepo := postgres2.NewUserRepository(db)
	orderRepo := postgres3.NewOrderRepository(db)
	historyRepo := postgres4.NewHistoryRepository(db)

	return &App{
		authUC: usecase.NewAuthUseCase(
			userRepo,
			"hash_salt",
			[]byte("signing_key"),
			time.Duration(86400),
		),
		orderUC:   usecase2.NewOrderUseCase(orderRepo),
		historyUC: usecase3.NewHistoryUseCase(historyRepo),
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
	delivery.RegisterHTTPEndpoints(router, a.authUC, l)

	// API endpoints
	authMiddleware := delivery.NewAuthMiddleware(a.authUC, l)
	api := router.Group("/", authMiddleware)

	delivery2.RegisterHTTPEndpointsOrder(api, a.orderUC, l)
	delivery3.RegisterHTTPEndpoints(api, a.historyUC, l)

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
