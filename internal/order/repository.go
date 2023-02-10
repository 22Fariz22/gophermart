package order

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
)

//interface

type OrderRepository interface {
	PushOrder(ctx context.Context, number uint32) error                   //POST /api/user/orders
	GetOrders(ctx context.Context, number uint32) ([]entity.Order, error) //GET /api/user/orders
}
