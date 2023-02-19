package worker

import (
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"net/http"
)

type UseCase interface {
	CheckNewOrders(l logger.Interface) ([]*entity.Order, error)

	SendToAccrualBox(orders []*entity.Order, httpServer *http.Server) error
	SendToWaitListChannels()
	//SendToOrdersCannels(orders []*entity.Order) error

}
