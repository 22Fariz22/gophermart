package history

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
)

//interface

type HistoryRepository interface {
	GetBalance(ctx context.Context, l logger.Interface, user *entity.User) (*entity.User, error)
	Withdraw(ctx context.Context, l logger.Interface, user *entity.User, number string, withdraw int) error
	InfoWithdrawal(ctx context.Context, l logger.Interface, user *entity.User) ([]*entity.History, error)
}
