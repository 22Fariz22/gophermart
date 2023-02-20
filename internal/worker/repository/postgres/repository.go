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

	var arrRespAcc arrRespAccr
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

		arrRespAcc = append(arrRespAcc, respAcc)
	}

	return arrRespAcc, nil
}

func mockResponse(l logger.Interface, orders []*entity.Order) ([]*entity.History, error) {
	fmt.Println("mockResponse().")

	var arrResAcc arrRespAccr
	var respAcc respAccr

	for _, v := range orders {
		err := json.Unmarshal(v, &respAcc)
		if err != nil {
			l.Error("Unmarshal error: ", err)
		}

		arrResAcc = append(arrResAcc, respAcc)
	}

	return nil, nil
}

func (w *WorkerRepository) SendToWaitListChannels() {
	//TODO implement me
	panic("implement me")
}

//func (p *HTTPClientSetting) GetAccrualResponse(orderNumber string) (*AccrualSystemResp, error) {
//
//	url := p.accrualAddr + "/api/orders/" + orderNumber
//
//	fmt.Printf("URL REQUEST IS %s \n", url)
//
//	req, err := http.NewRequestWithContext(context.TODO(), "GET", url, http.NoBody)
//
//	if err != nil {
//		return nil, err
//	}
//
//	resp, err := p.httpClient.Do(req)
//
//	if err != nil {
//		fmt.Printf("Err send req %v \n", err)
//		return nil, err
//	}
//
//	defer resp.Body.Close()
//
//	body, err := io.ReadAll(resp.Body)
//
//	if err != nil {
//		fmt.Printf("Err read resp %v \n", err)
//
//		return nil, err
//	}
//
//	if resp.StatusCode == 429 {
//		fmt.Println("Too many request, retrying after 60 sec...")
//		return nil, ErrTooManyRequest
//	}
//
//	response := AccrualSystemResp{}
//
//	fmt.Printf("Body is %+v \n", body)
//
//	if len(body) == 0 {
//		fmt.Println("Empty response!")
//		fmt.Printf("Status Code is %d \n", resp.StatusCode)
//		fmt.Printf("Order id is %s \n", orderNumber)
//		return nil, ErrEmptyResponse
//	}
//
//	err = json.Unmarshal(body, &response)
//
//	if err != nil {
//		fmt.Printf("Err unmarshal resp %v \n", err)
//
//		return nil, err
//	}
//
//	fmt.Printf("RESP IS %+v \n", response)
//	return &response, nil
//}
