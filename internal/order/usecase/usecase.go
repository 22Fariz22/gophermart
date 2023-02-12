package usecase

import (
	"context"
	"fmt"
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

func (o *OrderUseCase) PushOrder(ctx context.Context, user *entity.User, number string) error {
	fmt.Println("order-uc-PushOrder().")
	eo := &entity.Order{
		ID:         "",
		UserID:     "",
		Number:     number,
		Status:     "NEW",
		UploadedAt: time.Now(),
	}
	return o.orderRepo.PushOrder(ctx, user, eo)
}

func (o *OrderUseCase) GetOrders(ctx context.Context, user *entity.User) ([]*entity.Order, error) {

	return []*entity.Order{}, nil
}
