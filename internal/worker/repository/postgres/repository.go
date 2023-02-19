package postgres

import (
	"context"
	"fmt"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"net/http"
)

type WorkerRepository struct {
	*postgres.Postgres
}

func NewWorkerRepository(db *postgres.Postgres) *WorkerRepository {
	return &WorkerRepository{db}
}

func (w *WorkerRepository) CheckNewOrders(l logger.Interface) ([]*entity.Order, error) {
	fmt.Println("in repo-CheckNewOrders()")

	ctx := context.Background()
	rows, err := w.Pool.Query(ctx, `SELECT number FROM orders
									WHERE order_status = 'NEW'`)
	if err != nil {
		l.Error("err in Pool.Query()", err)
		return nil, err
	}

	out := make([]*entity.Order, 0)

	for rows.Next() {
		order := new(entity.Order)
		err := rows.Scan(&order.Number)
		if err != nil {
			l.Error("err rows.Scan(): ", err)
			return nil, err
		}
		out = append(out, order)
		fmt.Println("out rows: ", out)
	}

	return out, nil
}

func (w *WorkerRepository) SendToAccrualBox(orders []*entity.Order, httpServer *http.Server) error {
	fmt.Println("in repo-SendToAccrualBox()")

	//requestURL := fmt.Sprintf("localhost:8081/api/orders/%d",number)
	//
	//http.NewRequest(http.MethodGet, "localhost:8081/api/orders/", nil)
	number := "12345678903"
	req, err := http.NewRequest("GET", "localhost:8081/api/orders/"+number, nil)
	if req != nil {
		fmt.Println("NewRequest() err:", err)
	}
	fmt.Println("req:", req)

	return nil
}

func (w *WorkerRepository) SendToWaitListChannels() {
	//TODO implement me
	panic("implement me")
}

//func (w *WorkerRepository) SendToOrdersCannels(orders []*entity.Order) error {
//	//TODO implement me
//	panic("implement me")
//}
