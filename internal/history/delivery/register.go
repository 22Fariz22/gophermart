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

	historyEndpoints := router.Group("/api/user")
	{
		historyEndpoints.GET("/balance", h.GetBalance)
		historyEndpoints.POST("/balance/withdraw", h.Withdraw)
		historyEndpoints.GET("/withdrawals", h.InfoWithdrawal)
	}
}
