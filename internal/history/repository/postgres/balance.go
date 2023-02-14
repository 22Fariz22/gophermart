package postgres

import (
	"context"
	"fmt"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/internal/history"
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
	fmt.Println("history-repo-GetBalance()-user: ", user)
	var ub UserBalance

	var u entity.User

	err := pgxscan.Get(ctx, b.Pool, &ub, `SELECT balance_total, withdraw_total
										FROM users
										where user_id = $1;`, user.ID)
	if err != nil {
		l.Error("history-repo-GetBalance()-err: ", err)
		return nil, err
	}
	fmt.Println("history-repo-GetBalance()-ub: ", ub)

	u.BalanceTotal = ub.Balance_total
	u.WithdrawTotal = ub.Withdraw_total

	return &u, nil
}

func (b *BalanceRepository) Withdraw(ctx context.Context, l logger.Interface, user *entity.User,
	number string, withdrawResp int) error {
	withdraw_total := 0

	// reflect.TypeOf(history.ErrNotEnoughFunds) :  *errors.errorString

	err := pgxscan.Get(ctx, b.Pool, &withdraw_total, `SELECT withdraw_total FROM users
									  WHERE user_id = $1`, user.ID)
	if err != nil {
		l.Error("history-repo-Withdraw()-err: ", err)
		return err
	}

	fmt.Println("history-repo-Withdraw()-withdraw_total: ", withdraw_total)
	if withdraw_total < withdrawResp || withdrawResp < 0 { //сравниваем наш баланс с запросом
		l.Error("history-repo-Withdraw()- withdraw_total<withdrawResp): ", history.ErrNotEnoughFunds)
		return history.ErrNotEnoughFunds
	}

	_, err = b.Pool.Exec(ctx, `UPDATE users SET withdraw_total = withdraw_total - $1
								WHERE user_id = $2`, withdrawResp, user.ID)
	if err != nil {
		l.Error("history-repo-Withdraw()-exec update: ", err)
		return err
	}
	return nil
}

func (b *BalanceRepository) InfoWithdrawal(ctx context.Context, l logger.Interface,
	user *entity.User) ([]entity.History, error) {
	//TODO implement me
	panic("implement me")
}

func NewBalanceRepository(db *postgres.Postgres) *BalanceRepository {
	return &BalanceRepository{db}
}
