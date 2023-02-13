package postgres

import (
	"context"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"github.com/jackc/pgx/v5"
)

type BalanceRepository struct {
	*postgres.Postgres
}

func (b *BalanceRepository) GetBalance(ctx context.Context, l logger.Interface, user *entity.User) (*entity.User, error) {
	var userBalance entity.User

	if err:=
}

func (b *BalanceRepository) Withdraw(ctx context.Context, number uint32, withdraw uint32) error {
	//TODO implement me
	panic("implement me")
}

func (b *BalanceRepository) InfoWithdrawal(ctx context.Context) ([]entity.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func NewBalanceRepository(db *postgres.Postgres) *BalanceRepository {
	return &BalanceRepository{db}
}
