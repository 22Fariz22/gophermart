package usecase

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/balance"
	"github.com/22Fariz22/gophermart/internal/entity"
)

type BalanceUseCase struct {
	BalanceRepo balance.BalanceRepository
}

func NewBalanceUseCase(balanceRepo balance.BalanceRepository) *BalanceUseCase {
	return &BalanceUseCase{
		BalanceRepo: balanceRepo,
	}
}

func (b *BalanceUseCase) GetBalance(ctx context.Context, user *entity.User) (uint32, error) {

	return 0, nil
}

func (b *BalanceUseCase) Withdraw(ctx context.Context, number uint32, withdraw uint32) error {

	return nil
}

func (b *BalanceUseCase) InfoWithdrawal(ctx context.Context) ([]entity.Balance, error) {

	return nil, nil
}
