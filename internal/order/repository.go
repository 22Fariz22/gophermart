package order

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
)

//interface

type OrderRepository interface {
	PushOrder(ctx context.Context, user *entity.User, eo *entity.Order) error
	GetOrders(ctx context.Context, user *entity.User) ([]*entity.Order, error)
}
