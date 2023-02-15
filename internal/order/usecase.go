package order

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
)

type UseCase interface {
	PushOrder(ctx context.Context, l logger.Interface, user *entity.User, number string) error
	GetOrders(ctx context.Context, l logger.Interface, user *entity.User) ([]*entity.Order, error)
}
