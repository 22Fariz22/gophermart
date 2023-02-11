package delivery

import (
	"github.com/22Fariz22/gophermart/internal/order"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpointsOrder(router *gin.RouterGroup, uc order.UseCase, l logger.Interface) {
	h := NewHandler(uc, l)

	orders := router.Group("/")
	{
		orders.POST("orders", h.PushOrder)
		orders.GET("orders", h.GetOrders)
	}
}
