package repository

import (
	"github.com/22Fariz22/gophermart/internal/worker"
	"github.com/22Fariz22/gophermart/pkg/postgres"
)

type WorkerRepository struct {
	*postgres.Postgres
}

func NewWorkerRepository(db *postgres.Postgres) *WorkerRepository {
	return &WorkerRepository{db}
}

func SendToAccrualBox(orders []worker.NewOrders) error {

	return nil
}
