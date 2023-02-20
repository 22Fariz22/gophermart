package worker

import (
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
)

type WorkerRepository interface {
	CheckNewOrders(l logger.Interface) ([]*entity.Order, error)
	SendToAccrualBox(l logger.Interface, orders []*entity.Order) ([]*entity.History, error)
	SendToWaitListChannels()
}
