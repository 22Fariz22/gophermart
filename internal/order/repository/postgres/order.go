package postgres

import (
	"context"
	"fmt"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/postgres"
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
	UploadedAt time.Time
}

func NewOrderRepository(db *postgres.Postgres) *OrderRepository {
	return &OrderRepository{db}
}

func (o *OrderRepository) PushOrder(ctx context.Context, user *entity.User, eo *entity.Order) error {
	//eo.UserID = user.ID
	fmt.Println("number: ", eo.Number)
	fmt.Println("eo: ", eo)
	fmt.Println("user: ", user)
	_, err := o.Pool.Exec(ctx, `INSERT INTO orders (user_id, number, order_status, uploaded_at)
								VALUES ($1,$2,$3,$4)`,
		user.ID, eo.Number, eo.Status, eo.UploadedAt)
	if err != nil {
		fmt.Println("db-order-PushOrder()-err(1): ", err)
		return err
	}
	return nil
}

func (o OrderRepository) GetOrders(ctx context.Context) ([]entity.Order, error) {
	//TODO implement me
	panic("implement me")
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
