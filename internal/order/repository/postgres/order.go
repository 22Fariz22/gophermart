package postgres

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/postgres"
)

type OrderRepository struct {
	*postgres.Postgres
}

func NewOrderRepository(db *postgres.Postgres) *OrderRepository {
	return &OrderRepository{db}
}

func (o OrderRepository) PushOrder(ctx context.Context, user *entity.User, eo *entity.Order) error {
	eo.UserID = user.ID

	return nil
}

func (o OrderRepository) GetOrders(ctx context.Context, number uint32) ([]entity.Order, error) {
	//TODO implement me
	panic("implement me")
}
