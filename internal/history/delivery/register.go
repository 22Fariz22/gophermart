package delivery

import (
	"fmt"
	"github.com/22Fariz22/gophermart/internal/history"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc history.UseCase, l logger.Interface) {
	fmt.Println("auth-register-RegisterHTTPEndpoints(): ")

	h := NewHandler(uc, l)

	balanceEndpoints := router.Group("/api/user")
	{
		balanceEndpoints.GET("/history", h.GetBalance)
		balanceEndpoints.POST("/history/withdraw", h.Withdraw)
		balanceEndpoints.GET("/history/withdrawals", h.InfoWithdrawal)
	}
}
