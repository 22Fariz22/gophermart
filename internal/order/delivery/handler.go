package delivery

import (
	"github.com/22Fariz22/gophermart/internal/order"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type Order struct {
	ID         uint32
	Number     uint32
	Status     string
	UploadedAt time.Time
}

type Handler struct {
	useCase order.UseCase
	l       logger.Interface
}

func NewHandler(useCase order.UseCase, l logger.Interface) *Handler {
	return &Handler{
		useCase: useCase,
		l:       l,
	}
}

func (h *Handler) PushOrder(c *gin.Context) {

}

func (h *Handler) GetOrders(c *gin.Context) {

}
