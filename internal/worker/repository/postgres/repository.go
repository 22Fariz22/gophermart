package postgres

import (
	"context"
	"fmt"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"io"
	"log"
	"net/http"
	"net/url"
	"path"
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

	number := "12345678903"
	//req, err := http.NewRequest("GET", "localhost:8081/api/orders/"+number, nil)
	//if req != nil {
	//	fmt.Println("NewRequest() err:", err)
	//}
	//fmt.Println("req:", req)

	AccrualSystemAddress := "http://127.0.0.1:8080"

	reqURL, err := url.Parse(AccrualSystemAddress)
	fmt.Println("url.Parse")
	if err != nil {
		log.Fatalln("Wrong accrual system URL:", err)
	}

	reqURL.Path = path.Join("/api/orders/", number)
	fmt.Println("path.Join")
	r, err := http.Get(reqURL.String())
	fmt.Println("http.Get")
	if err != nil {
		log.Println("Can't get order updates from external API:", err)
	}

	body, err := io.ReadAll(r.Body)
	fmt.Println("io.ReadAll(r.Body)")
	defer r.Body.Close()

	if err != nil {
		log.Println("Can't read response body:", err)
	}
	fmt.Println("body:", body)

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
