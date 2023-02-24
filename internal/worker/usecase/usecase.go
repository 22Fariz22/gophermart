package usecase

import (
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/internal/worker"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"log"
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
	log.Println("worker-uc-CheckNewOrders().")
	//return w.CheckNewOrders(l)
	return w.WorkerRepo.CheckNewOrders(l)

}

func (w *WorkerUseCase) SendToAccrualBox(l logger.Interface, orders []*entity.Order) ([]*entity.History, error) {
	log.Println("worker-uc-SendToAccrualBox().")
	return w.WorkerRepo.SendToAccrualBox(l, orders)
}

//func (w *WorkerUseCase) SendToWaitListChannels() {
//	//TODO implement me
//	panic("implement me")
//}
