package usecase

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/internal/history"
	"github.com/22Fariz22/gophermart/pkg/logger"
)

type HistoryUseCase struct {
	HistoryRepo history.HistoryRepository
}

func NewHistoryUseCase(balanceRepo history.HistoryRepository) *HistoryUseCase {
	return &HistoryUseCase{
		HistoryRepo: balanceRepo,
	}
}

func (b *HistoryUseCase) GetBalance(ctx context.Context, l logger.Interface, user *entity.User) (*entity.User, error) {
	return b.HistoryRepo.GetBalance(ctx, l, user)
}

func (b *HistoryUseCase) Withdraw(ctx context.Context, l logger.Interface, user *entity.User, number string, withdraw int) error {
	return b.HistoryRepo.Withdraw(ctx, l, user, number, withdraw)
}

func (b *HistoryUseCase) InfoWithdrawal(ctx context.Context, l logger.Interface,
	user *entity.User) ([]*entity.History, error) {
	return b.HistoryRepo.InfoWithdrawal(ctx, l, user)
}
