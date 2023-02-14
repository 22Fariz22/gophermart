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

	var u entity.User

	err := pgxscan.Get(ctx, b.Pool, &ub, `SELECT balance_total, withdraw_total
										FROM users
										where user_id = $1;`, user.ID)
	if err != nil {
		l.Error("balance-repo-GetBalance()-err: ", err)
		return nil, err
	}
	fmt.Println("balance-repo-GetBalance()-ub: ", ub)

	u.BalanceTotal = ub.Balance_total
	u.WithdrawTotal = ub.Withdraw_total

	return &u, nil
}

func (b *BalanceRepository) Withdraw(ctx context.Context, l logger.Interface, user *entity.User,
	number string, withdraw int) error {
	withdraw_total := 0

	err := pgxscan.Get(ctx, b.Pool, &withdraw_total, `SELECT withdraw_total FROM users
									  WHERE user_id = $1`, user.ID)
	if err != nil {
		l.Error("balance-repo-Withdraw()-err: ", err)
		return err
	}
	fmt.Println("balance-repo-Withdraw()-withdraw_total: ", withdraw_total)

	return nil
}

func (b *BalanceRepository) InfoWithdrawal(ctx context.Context) ([]entity.Balance, error) {
	//TODO implement me
	panic("implement me")
}

func NewBalanceRepository(db *postgres.Postgres) *BalanceRepository {
	return &BalanceRepository{db}
}
