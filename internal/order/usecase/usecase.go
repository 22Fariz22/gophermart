package usecase

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/internal/order"
)

type OrderUserCase struct {
	orderRepo order.OrderRepository
}

func NewOrderUserCase(orderRepo order.OrderRepository) *OrderUserCase {
	return &OrderUserCase{orderRepo: orderRepo}
}

func (o *OrderUserCase) PushOrder(ctx context.Context, number uint32) error {

	return nil
}

func (o *OrderUserCase) GetOrders(ctx context.Context) ([]entity.Order, error) {

	return []entity.Order{}, nil
}
