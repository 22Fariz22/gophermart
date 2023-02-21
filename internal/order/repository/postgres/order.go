package postgres

import (
	"context"
	"fmt"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/internal/order"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"github.com/georgysavva/scany/v2/pgxscan"
	"strconv"
	"time"
)

type OrderRepository struct {
	*postgres.Postgres
}

type Order struct {
	ID         string
	UserID     string
	Number     string
	Status     string
	Accrual    uint32
	UploadedAt time.Time
}

func NewOrderRepository(db *postgres.Postgres) *OrderRepository {
	return &OrderRepository{db}
}

type existOrder struct {
	uID    string `db:"user_id"`
	number string `db:"number"`
}

func (o *OrderRepository) PushOrder(ctx context.Context, l logger.Interface, user *entity.User, eo *entity.Order) error {
	var existUser int
	var eOrd existOrder

	fmt.Println("order-PushOrder()-number:", eo.Number)
	// поменять на другой запрос
	_ = o.Pool.QueryRow(ctx, `SELECT user_id FROM orders where number = $1;`, eo.Number).Scan(&existUser)

	if err := pgxscan.Get(ctx, o.Pool, &eOrd, `SELECT user_id, number FROM orders where number = $1;`,
		eo.ID, eo.Number); err != nil {
		l.Info("err in pgxsan.Get():", err)
	}
	fmt.Println("eOrd: ", eOrd)

	existUserConvToStr, err := strconv.Atoi(user.ID)
	if err != nil {
		l.Error("err in strconv.Atoi(user.ID)")
		return err
	}

	fmt.Println(" existUser == existUserConvToStr", existUser, existUserConvToStr)
	if existUser == existUserConvToStr {
		l.Info("Number Has Already Been Uploaded")
		return order.ErrNumberHasAlreadyBeenUploaded
	}
	if existUser != existUserConvToStr {
		l.Info("Number Has Already Been Uploaded By AnotherUser")
		return order.ErrNumberHasAlreadyBeenUploadedByAnotherUser
	}

	tx, err := o.Pool.Begin(ctx)
	if err != nil {
		l.Error("tx err: ", err)
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `INSERT INTO orders (user_id, number, order_status, uploaded_at)
								VALUES ($1,$2,$3,$4)`,
		user.ID, eo.Number, eo.Status, eo.UploadedAt)
	if err != nil {
		fmt.Println("db-order-PushOrder()-err(1): ", err)
		return err
	}

	_, err = tx.Exec(ctx, `INSERT INTO history (user_id, number, order_status, uploaded_at)
								VALUES ($1,$2,$3,$4)`,
		user.ID, eo.Number, eo.Status, eo.UploadedAt)
	if err != nil {
		fmt.Println("db-order-PushOrder()-err(1): ", err)
		return err
	}
	err = tx.Commit(ctx)
	if err != nil {
		l.Error("commit err: ", err)
		return err
	}

	return nil
}

func (o OrderRepository) GetOrders(ctx context.Context, l logger.Interface, user *entity.User) ([]*entity.Order, error) {
	rows, err := o.Pool.Query(ctx, `SELECT order_id, number, order_status, accrual, uploaded_at FROM orders
									WHERE user_id = $1`, user.ID)
	if err != nil {
		l.Error("err Pool.Query: ", err)
		return nil, err
	}

	out := make([]*entity.Order, 0)

	for rows.Next() {
		order := new(entity.Order)
		err := rows.Scan(&order.ID, &order.Number, &order.Status, &order.Accrual, &order.UploadedAt)
		if err != nil {
			l.Error("err rows.Scan(): ", err)
			return nil, err
		}
		out = append(out, order)
	}

	return out, nil
}

func toModel(o *entity.Order) *Order {
	return &Order{
		ID:         o.ID,
		UserID:     o.UserID,
		Number:     o.Number,
		Status:     o.Status,
		UploadedAt: o.UploadedAt,
	}
}
