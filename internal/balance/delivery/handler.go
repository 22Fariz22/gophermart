package delivery

import (
	"github.com/22Fariz22/gophermart/internal/balance"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/gin-gonic/gin"
	"time"
)

type Balance struct {
	ID           uint32
	OrderID      uint32
	UserID       uint32
	Accrual      uint32
	Status       uint8
	UploadedAt   time.Time
	WithdrawDate time.Time
}

type Handler struct {
	useCase balance.UseCase
	l       logger.Interface
}

func NewHandler(useCase balance.UseCase, l logger.Interface) *Handler {
	return &Handler{
		useCase: useCase,
		l:       l,
	}
}

func (h *Handler) GetBalance(c *gin.Context) {

}

func (h *Handler) InfoWithdrawal(c *gin.Context) {

}

func (h *Handler) Withdraw(c *gin.Context) {

}
