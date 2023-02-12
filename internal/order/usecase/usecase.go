package usecase

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/internal/order"
	"time"
)

type OrderUseCase struct {
	orderRepo order.OrderRepository
}

func NewOrderUseCase(orderRepo order.OrderRepository) *OrderUseCase {
	return &OrderUseCase{orderRepo: orderRepo}
}

func (o *OrderUseCase) PushOrder(ctx context.Context, user *entity.User, number uint32) error {
	eo := &entity.Order{
		ID:         "",
		UserID:     "",
		Number:     number,
		Status:     "",
		UploadedAt: time.Now(),
	}
	return o.orderRepo.PushOrder(ctx, user, eo)
}

func (o *OrderUseCase) GetOrders(ctx context.Context) ([]entity.Order, error) {

	return []entity.Order{}, nil
}
