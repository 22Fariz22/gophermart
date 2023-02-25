package postgres

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/22Fariz22/gophermart/internal/config"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"github.com/22Fariz22/gophermart/pkg/postgres"
	"io"
	"log"
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
	log.Println("worker-repo-CheckNewOrders()")

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
	Order   string  `json:"order"`
	Status  string  `json:"status"`
	Accrual float64 `json:"accrual"`
}

//SendToAccrualBox отправляем запрос accrual system и возвращаем ответ от него
func (w *WorkerRepository) SendToAccrualBox(l logger.Interface, cfg *config.Config, orders []*entity.Order) ([]*entity.History, error) {
	log.Println("worker-repo-SendToAccrualBox()")

	//структура json ответа от accrual sysytem
	var resAccrSys ResAccrualSystem

	// считываем из env переменную ACCRUAL_SYSTEM_ADDRESS
	accrualSystemAddress := cfg.AccrualSystemAddress

	//возвращаем мок, если запускаем приложение у себя локально
	//if accrualSystemAddress == "mock" {
	//	return mockResponse(l, orders)
	//}

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

		fmt.Println("body from response accrual system:: ", string(body))

		// if status == 204: do update set order_status = INVALID, history_status = INVALID
		if r.StatusCode == 204 {
			// do update in data in tables orders and history
			if err := update(w, l, ResAccrualSystem{
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
			update(w, l, resAccrSys)
		}

		if r.StatusCode == 429 {
			sleep, err := time.ParseDuration(r.Header.Get("Retry-After"))
			if err != nil {
				l.Error("worker-repo-SendToAccrualBox()-status 429- time.ParseDuration()-err: ", err)
				time.Sleep(60 * time.Second)
			}
			time.Sleep(sleep)
		}

		if r.StatusCode == 500 {
			l.Error("worker-repo-SendToAccrualBox()-status500.")
			return nil, err
		}
	}

	return nil, nil
}

func update(w *WorkerRepository, l logger.Interface, resAcc ResAccrualSystem) error {
	log.Println("worker-repo-updateWithStatus()")

	ctx := context.Background()

	//UPDATE в таблице History и Orders
	log.Println("worker-repo-updateWithStatus()- start begin tx.")
	tx, err := w.Pool.Begin(ctx)
	if err != nil {
		l.Error("tx err: ", err)
		return err
	}
	defer tx.Rollback(ctx)

	_, err = tx.Exec(ctx, `UPDATE orders SET order_status =  $1, accrual = $2
							where number = $3`, resAcc.Status, resAcc.Accrual*100, resAcc.Order)
	if err != nil {
		l.Error("error in Exec UPDATE: ", err)
		return err
	}

	// надо ли делать UPDATE в history???

	log.Println("worker-repo-updateWithStatus().-tx commit.")

	err = tx.Commit(ctx)
	if err != nil {
		l.Error("worker-repo-updateWithStatus() -tx.commit err: ", err)
		return err
	}

	return nil
}

//func (w *WorkerRepository) SendToWaitListChannels() {
//	//TODO implement me
//	panic("implement me")
//}

//func checkStatus(w *WorkerRepository, l logger.Interface, resAcc ResAccrualSystem) error {
//	log.Println("worker-repo-checkStatus()")
//	err := updateWithStatus(w, l, resAcc)
//	if err != nil {
//		l.Error("worker-repo-checkStatus()-updateWithStatus()-err",err)
//		return err
//	}
//	return nil
//}
