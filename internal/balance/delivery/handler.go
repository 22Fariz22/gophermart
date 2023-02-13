package delivery

import (
	"fmt"
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/22Fariz22/gophermart/internal/balance"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
)

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

type BalanceResponce struct {
	current   uint32 `json:"current"`
	withdrawn uint32 `json:"withdrawn"`
}

func (h *Handler) GetBalance(c *gin.Context) {
	fmt.Println("balance-handler-GetBalance().")
	user := c.MustGet(auth.CtxUserKey).(*entity.User)
	fmt.Println("balance-handler-GetBalance()-user: ", user)

	balance, err := h.useCase.GetBalance(c.Request.Context(), h.l, user)
	if err != nil {
		h.l.Error("Status Internal ServerError: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	fmt.Println("balance-handler-GetBalance()-balance: ", balance)

}

func (h *Handler) InfoWithdrawal(c *gin.Context) {

}

func (h *Handler) Withdraw(c *gin.Context) {

}
