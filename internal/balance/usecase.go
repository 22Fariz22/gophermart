package balance

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
)

type UseCase interface {
	GetBalance(ctx context.Context, user *entity.User) (uint32, error)  //GET /api/user/balance
	Withdraw(ctx context.Context, number uint32, withdraw uint32) error //POST /api/user/balance/withdraw
	InfoWithdrawal(ctx context.Context) ([]entity.Balance, error)       //GET /api/user/withdrawals
}
