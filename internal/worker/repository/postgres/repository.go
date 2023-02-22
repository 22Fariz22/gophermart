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
	"time"
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
									WHERE order_status IN( 'NEW','PROCESSING')`)
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

//status 200->PROCESSED,  204->INVALID,  401,  429, 500

// структура json ответа от accrual sysytem
type ResAccrualSystem struct {
	Order   string `json:"order"`
	Status  string `json:"status"`
	Accrual int    `json:"accrual"`
}

//SendToAccrualBox отправляем запрос accrual system и возвращаем ответ от него
func (w *WorkerRepository) SendToAccrualBox(l logger.Interface, orders []*entity.Order) ([]*entity.History, error) {
	fmt.Println("in repo-SendToAccrualBox()")
	//var arrResAcc arrRespAccr

	//структура json ответа от accrual sysytem
	var resAccrSys ResAccrualSystem

	// считываем из env переменную ACCRUAL_SYSTEM_ADDRESS
	accrualSystemAddress := viper.GetString("r")

	//возвращаем мок, если запускаем приложение у себя локально
	if accrualSystemAddress == "mock" {
		return mockResponse(l, orders)
	}

	reqURL, err := url.Parse(accrualSystemAddress)
	fmt.Println("url.Parse")
	if err != nil {
		l.Error("incorrect ACCRUAL_SYSTEM_ADDRESS:", err)
		return nil, err // выходим если адрес accrual system некорректный
	}

	// проходимся по списку ордеров и обращаемся к accrual system
	for _, v := range orders {
		reqURL.Path = path.Join("/api/orders/", v.Number)
		fmt.Println("reqURL.String()", reqURL.String())

		r, err := http.Get(reqURL.String())
		fmt.Println("http.Get")
		if err != nil {
			l.Error("can't do request: ", err)
			return nil, err //выходим из цикла, если не получился запрос к accrual system
		}

		body, err := io.ReadAll(r.Body)
		fmt.Println("io.ReadAll(r.Body)")
		defer r.Body.Close()
		if err != nil {
			l.Error("Can't read response body: ", err)
			continue //переходим к следущей итерации
		}

		fmt.Println("body: ", string(body))

		// if status == 204: do update set order_status = INVALID, history_status = INVALID
		if r.StatusCode == 204 {
			// do update in data in tables orders and history
			if err := checkStatus(w, l, ResAccrualSystem{
				Order:   v.Number,
				Status:  "INVALID",
				Accrual: 0,
			}); err != nil {
				return nil, err // определить какой error
			}
		}

		// if status == 200: делаем update in db to PROCESSED
		if r.StatusCode == 200 {
			// do  unmarshall
			err = json.Unmarshal(body, &resAccrSys)
			if err != nil {
				l.Error("Unmarshal error: ", err)
			}

			//do update
			checkStatus(w, l, resAccrSys)
		}

		if r.StatusCode == 429 {
			sleep, err := time.ParseDuration(r.Header.Get("Retry-After"))
			if err != nil {
				time.Sleep(60 * time.Second)
			}
			time.Sleep(sleep)
		}

		if r.StatusCode == 500 {
			return nil, err
		}

	}

	return nil, nil
}

func checkStatus(w *WorkerRepository, l logger.Interface, resAcc ResAccrualSystem) error {
	err := updateWithStatus(w, l, resAcc)
	if err != nil {
		return err
	}
	return nil
}

func updateWithStatus(w *WorkerRepository, l logger.Interface, resAcc ResAccrualSystem) error {
	ctx := context.Background()

	//UPDATE в таблице History и Orders
	_, err := w.Pool.Exec(ctx, `UPDATE orders SET order_status =  $1, accrual = $2
							where number = $3`, resAcc.Status, resAcc.Accrual, resAcc.Order)
	if err != nil {
		l.Error("error in Exec UPDATE: ", err)
		return err
	}

	return nil
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

//func (w *WorkerRepository) SendToWaitListChannels() {
//	//TODO implement me
//	panic("implement me")
//}
