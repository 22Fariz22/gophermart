package postgres

import (
	"context"
	"fmt"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"log"
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

func (o *OrderRepository) PushOrder(ctx context.Context, l logger.Interface, user *entity.User, eo *entity.Order) error {
	_, err := o.Pool.Exec(ctx, `INSERT INTO orders (user_id, number, order_status, uploaded_at)
								VALUES ($1,$2,$3,$4)`,
		user.ID, eo.Number, eo.Status, eo.UploadedAt)
	if err != nil {
		fmt.Println("db-order-PushOrder()-err(1): ", err)
		return err
	}
	return nil
}

func (o OrderRepository) GetOrders(ctx context.Context, l logger.Interface, user *entity.User) ([]*entity.Order, error) {
	fmt.Println("order-repo-GetOrders().")
	rows, err := o.Pool.Query(ctx, `SELECT order_id, number, order_status, accrual, uploaded_at FROM orders
									WHERE user_id = $1`, user.ID)
	if err != nil {
		return nil, err
	}

	out := make([]*entity.Order, 0)

	for rows.Next() {
		order := new(entity.Order)
		err := rows.Scan(&order.ID, &order.Number, &order.Status, &order.Accrual, &order.UploadedAt)
		if err != nil {
			log.Println("order-repo-GetOrders()-rows.Scan()-err: ", err)
			return nil, err
		}
		fmt.Println("order-repo-GetOrders()-order: ", order)
		out = append(out, order)
	}

	fmt.Println("order-repo-GetOrders()-out: ", out)
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
