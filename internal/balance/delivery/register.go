package delivery

import (
	"fmt"
	"github.com/22Fariz22/gophermart/internal/balance"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpoints(router *gin.RouterGroup, uc balance.UseCase, l logger.Interface) {
	fmt.Println("auth-register-RegisterHTTPEndpoints(): ")

	h := NewHandler(uc, l)

	balanceEndpoints := router.Group("/api/user")
	{
		balanceEndpoints.GET("/balance", h.GetBalance)
		balanceEndpoints.POST("/balance/withdraw", h.Withdraw)
		balanceEndpoints.GET("/balance/withdrawals", h.InfoWithdrawal)
	}
}
