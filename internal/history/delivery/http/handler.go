package http

import (
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/internal/history"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Handler struct {
	useCase history.UseCase
	l       logger.Interface
}

func NewHandler(useCase history.UseCase, l logger.Interface) *Handler {
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
	user := c.MustGet(auth.CtxUserKey).(*entity.User)

	u, err := h.useCase.GetBalance(c.Request.Context(), h.l, user)
	if err != nil {
		h.l.Error("Status Internal ServerError: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	br := toBalanceResponce(u)

	c.JSON(http.StatusOK, BalanceResponce{
		Current:   br.Current,
		Withdrawn: br.Withdrawn,
	})
}

func toBalanceResponce(u *entity.User) *BalanceResponce {
	return &BalanceResponce{
		Current:   u.BalanceTotal,
		Withdrawn: u.WithdrawTotal,
	}
}

type InputWithdraw struct {
	Order string `json:"order"`
	Sum   int    `json:"sum"`
}

func (h *Handler) Withdraw(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*entity.User)

	inp := new(InputWithdraw)
	if err := c.BindJSON(inp); err != nil {
		h.l.Error("history-handler-Withdraw()-BindJSON-err: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err := h.useCase.Withdraw(c.Request.Context(), h.l, user, inp.Order, inp.Sum)
	if err != nil {
		if err == history.ErrNotEnoughFunds { //если не достаточно баллов
			h.l.Info("Not Enough Funds")
			c.AbortWithStatus(http.StatusPaymentRequired)
		}

		h.l.Error("Status Internal Server Error: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	// добавить проверку алгоритма Луна, при ошибке который верет 422

	c.Status(http.StatusOK)
}

//type HistoryResponse struct {
//	HistoryResp []*entity.History `json:"history_resp"`
//}
type HistoryResp struct {
	Number      string    `json:"number"`
	Sum         int       `json:"sum"`
	ProcessedAt time.Time `json:"processed_at"`
}

func (h *Handler) InfoWithdrawal(c *gin.Context) {
	user := c.MustGet(auth.CtxUserKey).(*entity.User)

	hist, err := h.useCase.InfoWithdrawal(c.Request.Context(), h.l, user)
	if err != nil {
		if err == history.ErrThereIsNoWithdrawal { //нет списаний
			h.l.Error("There Is No Withdrawal")
			c.AbortWithStatus(http.StatusNoContent)
		}
		h.l.Error("")
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, hist)
}
