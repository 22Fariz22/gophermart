package order

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
)

type UseCase interface {
	PushOrder(ctx context.Context, number uint32) error
	GetOrders(ctx context.Context) ([]entity.Order, error)
}
