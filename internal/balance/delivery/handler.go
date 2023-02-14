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
	Current   int `json:"current"`
	Withdrawn int `json:"withdrawn"`
}

func (h *Handler) GetBalance(c *gin.Context) {
	fmt.Println("balance-handler-GetBalance().")
	user := c.MustGet(auth.CtxUserKey).(*entity.User)
	fmt.Println("balance-handler-GetBalance()-user: ", user)

	u, err := h.useCase.GetBalance(c.Request.Context(), h.l, user)
	if err != nil {
		h.l.Error("Status Internal ServerError: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	fmt.Println("balance-handler-GetBalance()-balance: ", u)
	br := toBalanceResponce(u)

	fmt.Println("balance-handler-GetBalance()-br:", br)

	c.JSON(http.StatusOK, BalanceResponce{
		Current:   br.Current,
		Withdrawn: br.Withdrawn,
	})
}
func toBalanceResponce(u *entity.User) *BalanceResponce {
	return &BalanceResponce{
		Current:   u.BalanceTotal / 100,
		Withdrawn: u.WithdrawTotal,
	}
}

func (h *Handler) InfoWithdrawal(c *gin.Context) {

}

func (h *Handler) Withdraw(c *gin.Context) {

}
