package usecase

import (
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

func (w *WorkerUseCase) SendToAccrualBox(orders []worker.NewOrders) error {

	return nil
}