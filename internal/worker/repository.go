package worker

import "github.com/22Fariz22/gophermart/internal/entity"

type WorkerRepository interface {
	CheckNewOrders() ([]*entity.Order, error)
	SendToOrdersCannels(orders []*entity.Order) error
	SendToAccrualBox(orders []NewOrders) error
	SendToWaitListChannels()
}
