package usecase

import (
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/internal/worker"
)

type WorkerUseCase struct {
	WorkerRepo worker.WorkerRepository
}

func NewWorkerUseCase(repo worker.WorkerRepository) *WorkerUseCase {
	return &WorkerUseCase{
		WorkerRepo: repo,
	}
}

func (w *WorkerUseCase) CheckNewOrders() ([]*entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (w *WorkerUseCase) SendToOrdersCannels(orders []*entity.Order) error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkerUseCase) SendToWaitListChannels() {
	//TODO implement me
	panic("implement me")
}

func (w *WorkerUseCase) SendToAccrualBox(orders []worker.NewOrders) error {

	return nil
}
