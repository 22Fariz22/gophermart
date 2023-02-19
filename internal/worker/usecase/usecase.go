package usecase

import (
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/internal/worker"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"net/http"
)

type WorkerUseCase struct {
	WorkerRepo worker.WorkerRepository
}

func NewWorkerUseCase(repo worker.WorkerRepository) *WorkerUseCase {
	return &WorkerUseCase{
		WorkerRepo: repo,
	}
}

func (w *WorkerUseCase) CheckNewOrders(l logger.Interface) ([]*entity.Order, error) {
	return w.CheckNewOrders(l)
}

func (w *WorkerUseCase) SendToAccrualBox(orders []*entity.Order, httpServer *http.Server) error {
	return w.SendToAccrualBox(orders, httpServer)
}

func (w *WorkerUseCase) SendToWaitListChannels() {
	//TODO implement me
	panic("implement me")
}

//func (w *WorkerUseCase) SendToOrdersCannels(orders []*entity.Order) error {
//	//TODO implement me
//	panic("implement me")
//}
