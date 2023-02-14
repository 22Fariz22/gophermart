package usecase

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/balance"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
)

type BalanceUseCase struct {
	BalanceRepo balance.BalanceRepository
}

func NewBalanceUseCase(balanceRepo balance.BalanceRepository) *BalanceUseCase {
	return &BalanceUseCase{
		BalanceRepo: balanceRepo,
	}
}

func (b *BalanceUseCase) GetBalance(ctx context.Context, l logger.Interface, user *entity.User) (*entity.User, error) {
	u, err := b.BalanceRepo.GetBalance(ctx, l, user)
	if err != nil {
		l.Error("balance-uc-GetBalance()-err: ", err)
		return nil, err
	}
	return u, nil
}

func (b *BalanceUseCase) Withdraw(ctx context.Context, number uint32, withdraw uint32) error {
	//TODO implement me
	panic("implement me")
}

func (b *BalanceUseCase) InfoWithdrawal(ctx context.Context) ([]entity.Balance, error) {
	//TODO implement me
	panic("implement me")
}
