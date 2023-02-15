package delivery

import (
	"fmt"
	"github.com/22Fariz22/gophermart/internal/auth"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/internal/order"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
	"time"
)

type Order struct {
	ID         string    `json:"id"`
	Number     string    `json:"number"`
	Status     string    `json:"status"`
	Accrual    uint32    `json:"accrual"`
	UploadedAt time.Time `json:"uploaded_at"`
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

type Number struct {
	number string
}

func (h *Handler) PushOrder(c *gin.Context) {
	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		h.l.Error("Status Bad Request: ", err)
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Println("order-handler-PushOrder()-payload: ", string(payload))
	// еще добавить проверку Луна и к нему статус 422

	user := c.MustGet(auth.CtxUserKey).(*entity.User)

	if err := h.useCase.PushOrder(c.Request.Context(), h.l, user, string(payload)); err != nil {
		h.l.Error("Status Internal ServerError: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

type ordersResponse struct {
	Orders []*Order `json:"orders"`
}

func (h *Handler) GetOrders(c *gin.Context) {
	fmt.Println("order-handler-GetOrder()")

	user := c.MustGet(auth.CtxUserKey).(*entity.User)

	orders, err := h.useCase.GetOrders(c.Request.Context(), h.l, user)
	fmt.Println("order-handler-GetOrder()-orders: ", orders)
	if err != nil {
		h.l.Error("Status Internal ServerError: ", err)
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, &ordersResponse{
		Orders: toOrders(orders),
	})
}

func toOrders(os []*entity.Order) []*Order {
	out := make([]*Order, len(os))

	for i, o := range os {
		out[i] = toOrder(o)
	}
	return out
}

func toOrder(o *entity.Order) *Order {
	return &Order{
		ID:         o.ID,
		Number:     o.Number,
		Status:     o.Status,
		UploadedAt: time.Time{},
	}
}
