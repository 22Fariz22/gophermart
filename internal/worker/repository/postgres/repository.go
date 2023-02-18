package postgres

import (
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/internal/worker"
	"github.com/22Fariz22/gophermart/pkg/postgres"
)

type WorkerRepository struct {
	*postgres.Postgres
}

func NewWorkerRepository(db *postgres.Postgres) *WorkerRepository {
	return &WorkerRepository{db}
}

func (w *WorkerRepository) CheckNewOrders() ([]*entity.Order, error) {
	//TODO implement me
	panic("implement me")
}

func (w *WorkerRepository) SendToOrdersCannels(orders []*entity.Order) error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkerRepository) SendToAccrualBox(orders []worker.NewOrders) error {
	//TODO implement me
	panic("implement me")
}

func (w *WorkerRepository) SendToWaitListChannels() {
	//TODO implement me
	panic("implement me")
}
