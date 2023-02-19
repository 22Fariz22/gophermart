package usecase

import (
	"fmt"
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
	fmt.Println("in uc-CheckNewOrders()")
	//return w.CheckNewOrders(l)
	return w.WorkerRepo.CheckNewOrders(l)

}

func (w *WorkerUseCase) SendToAccrualBox(orders []*entity.Order, httpServer *http.Server) error {
	fmt.Println("in uc-SendToAccrualBox()")
	return w.WorkerRepo.SendToAccrualBox(orders, httpServer)
}

func (w *WorkerUseCase) SendToWaitListChannels() {
	//TODO implement me
	panic("implement me")
}
