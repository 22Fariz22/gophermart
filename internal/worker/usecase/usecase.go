package usecase

import (
	"fmt"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/internal/worker"
	"github.com/22Fariz22/gophermart/pkg/logger"
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

func (w *WorkerUseCase) SendToAccrualBox(l logger.Interface, orders []*entity.Order) ([]*entity.History, error) {
	fmt.Println("in uc-SendToAccrualBox()")
	return w.WorkerRepo.SendToAccrualBox(l, orders)
}

func (w *WorkerUseCase) SendToWaitListChannels() {
	//TODO implement me
	panic("implement me")
}
