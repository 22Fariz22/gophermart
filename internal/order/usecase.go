package order

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
)

type UseCase interface {
	PushOrder(ctx context.Context, user *entity.User, number string) error
	GetOrders(ctx context.Context, user *entity.User) ([]*entity.Order, error)
}
