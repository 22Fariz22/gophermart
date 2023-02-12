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

type Number struct {
	number string
}

func (h *Handler) PushOrder(c *gin.Context) {
	payload, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}
	fmt.Println("order-handler-PushOrder()-payload: ", string(payload))
	// еще добавить проверку Луна и к нему статус 422

	user := c.MustGet(auth.CtxUserKey).(*entity.User)
	if err := h.useCase.PushOrder(c.Request.Context(), user, string(payload)); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) GetOrders(c *gin.Context) {

}
