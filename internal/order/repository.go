package order

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
)

//interface

type OrderRepository interface {
	PushOrder(ctx context.Context, l logger.Interface, user *entity.User, eo *entity.Order) error
	GetOrders(ctx context.Context, l logger.Interface, user *entity.User) ([]*entity.Order, error)
}
