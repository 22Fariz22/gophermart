package usecase

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/internal/order"
)

type OrderUseCase struct {
	orderRepo order.OrderRepository
}

func NewOrderUseCase(orderRepo order.OrderRepository) *OrderUseCase {
	return &OrderUseCase{orderRepo: orderRepo}
}

func (o *OrderUseCase) PushOrder(ctx context.Context, number uint32) error {

	return nil
}

func (o *OrderUseCase) GetOrders(ctx context.Context) ([]entity.Order, error) {

	return []entity.Order{}, nil
}
