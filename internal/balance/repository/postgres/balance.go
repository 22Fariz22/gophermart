package postgres

import (
	"context"
	"fmt"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"github.com/georgysavva/scany/v2/pgxscan"
)

type BalanceRepository struct {
	*postgres.Postgres
}
type UserBalance struct {
	Balance_total  int
	Withdraw_total int
}

func (b *BalanceRepository) GetBalance(ctx context.Context, l logger.Interface, user *entity.User) (*entity.User, error) {
	fmt.Println("balance-repo-GetBalance()-user: ", user)
	var ub UserBalance
	//pgxscan.Select(ctx, b.Pool, &ub, `SELECT balance_total, withdraw_total FROM users
	//				where user_id = 1;`)
	err := pgxscan.Get(ctx, b.Pool, &ub, `SELECT balance_total, withdraw_total
										FROM users
										where user_id = $1;`, user.ID)
	if err != nil {
		l.Error("balance-repo-GetBalance()-err: ", err)
		return nil, err
	}
	fmt.Println("balance-repo-GetBalance()-ub: ", ub)
	var u entity.User
	u.BalanceTotal = ub.Balance_total
	u.WithdrawTotal = ub.Withdraw_total
	//fmt.Println(u.BalanceTotal)
	//fmt.Println(u.WithdrawTotal)
	return &u, nil
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
