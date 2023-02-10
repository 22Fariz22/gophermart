package delivery

import (
	"github.com/22Fariz22/gophermart/internal/order"
	"github.com/gin-gonic/gin"
)

func RegisterHTTPEndpointsOrder(router *gin.RouterGroup, uc order.UseCase) {
	h := NewHandler(uc)

	orders := router.Group("/")
	{
		orders.POST("orders", h.PushOrder)
		orders.GET("orders", h.GetOrders)
	}
}
