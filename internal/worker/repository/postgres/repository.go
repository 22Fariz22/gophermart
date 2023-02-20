package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"github.com/spf13/viper"
	"io"
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

type arrRespAccr []*entity.History // структура ответа от accrual system
type respAccr *entity.History

//SendToAccrualBox отправляем запрос accrual system и возвращаем ответ от него
func (w *WorkerRepository) SendToAccrualBox(l logger.Interface, orders []*entity.Order) ([]*entity.History, error) {
	fmt.Println("in repo-SendToAccrualBox()")

	var arrResAcc arrRespAccr
	var respAcc respAccr

	accrualSystemAddress := viper.GetString("r")

	//возвращаем мок, если запускаем приложение у себя локально
	if accrualSystemAddress == "mock" {
		return mockResponse(l, orders)
	}

	for _, v := range orders {
		reqURL, err := url.Parse(accrualSystemAddress) // считываем из env переменную ACCRUAL_SYSTEM_ADDRESS
		fmt.Println("url.Parse")
		if err != nil {
			l.Error("incorrect ACCRUAL_SYSTEM_ADDRESS:", err)
			return nil, err // выходим из цикла, если адрес accrual system некорректный
		}

		reqURL.Path = path.Join("/api/orders/", v.Number)
		fmt.Println("path.Join")

		r, err := http.Get(reqURL.String())
		fmt.Println("http.Get")
		if err != nil {
			l.Error("can't do request: ", err)
			//что возвращаем? выходим из цикла?
		}

		body, err := io.ReadAll(r.Body)
		fmt.Println("io.ReadAll(r.Body)")
		defer r.Body.Close()

		if err != nil {
			l.Error("Can't read response body: ", err)
			//что возвращаем? выходим из цикла?
		}
		fmt.Println("body: ", string(body))

		// unmarshall
		err = json.Unmarshal(body, &respAcc)
		if err != nil {
			l.Error("Unmarshal error: ", err)
		}

		arrResAcc = append(arrResAcc, respAcc)
	}

	return arrResAcc, nil
}

func mockResponse(l logger.Interface, orders []*entity.Order) ([]*entity.History, error) {
	fmt.Println("mockResponse().")

	var arrRA arrRespAccr

	for _, v := range orders {
		fmt.Println("range v in orders: ", v)
		respAcc := entity.History{
			ID:     v.ID,
			UserID: v.UserID,
			Number: v.Number,
			Status: "PROCESSED",
			Sum:    777,
		}
		arrRA = append(arrRA, &respAcc)
	}

	fmt.Println("ArrRA: ", arrRA)
	return arrRA, nil
}

func (w *WorkerRepository) SendToWaitListChannels() {
	//TODO implement me
	panic("implement me")
}
