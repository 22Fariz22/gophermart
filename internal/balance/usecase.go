package balance

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
)

type UseCase interface {
	GetBalance(ctx context.Context, l logger.Interface, user *entity.User) (*entity.User, error) //GET /api/user/balance
	Withdraw(ctx context.Context, number uint32, withdraw uint32) error                          //POST /api/user/balance/withdraw
	InfoWithdrawal(ctx context.Context) ([]entity.Balance, error)                                //GET /api/user/withdrawals
}
