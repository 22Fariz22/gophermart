package worker

import (
	"errors"
	"fmt"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"net/http"
	"sync"
	"time"
)

type Pool struct {
	httpServer *http.Server
	l          logger.Interface
	wg         sync.WaitGroup
	once       sync.Once
	shutDown   chan struct{}
	mainCh     chan workerData
	repository UseCase
}

func NewWorkerPool(repo UseCase, l logger.Interface, httpServer *http.Server) *Pool {
	return &Pool{
		httpServer: httpServer,
		l:          l,
		wg:         sync.WaitGroup{},
		once:       sync.Once{},
		shutDown:   make(chan struct{}),
		mainCh:     make(chan workerData, 10),
		repository: repo,
	}
}

type workerData struct {
	orders []*entity.Order
}

//  функция которая каждые 2 мин забирает из таблицы ордеры со статусом NEW и кладет их в каналы
func CollectNewOrders(uc UseCase, l logger.Interface, httpServer *http.Server) []*entity.Order {
	workers := NewWorkerPool(uc, l, httpServer)

	for {
		workers.RunWorkers(5)
		defer workers.Stop()

		time.Sleep(4 * time.Second)

		newOrders, err := workers.repository.CheckNewOrders(l) //получаем список новых ордеров
		fmt.Println("newOrders: ", newOrders)
		if err != nil {
			l.Info("err in CheckNewOrders(): ", err)
		}

		workers.AddJob(newOrders)
	}

}

type NewOrders struct {
	number string
}

func (w *Pool) AddJob(arr []*entity.Order) error {
	select {
	case <-w.shutDown:
		return errors.New("all channels are closed")
	case w.mainCh <- workerData{
		orders: arr,
	}:
		return nil
	}
}

// далее этот список передаем воркеру
func (w *Pool) RunWorkers(count int) {
	fmt.Println("start RunWorkers()")
	for i := 0; i < count; i++ {
		fmt.Println("start RunWorkers() for... count")
		w.wg.Add(1)
		go func() {
			defer w.wg.Done()
			for {
				select {
				case <-w.shutDown:
					fmt.Println("case <-w.shutDown.")
					//w.l.Info("channels are shutdown.")
					return
				case orders, ok := <-w.mainCh:
					fmt.Println("case <-w.mainCh.")
					if !ok {
						fmt.Println("case <-w.mainCh !ok")
						return
					}
					fmt.Println("SendToAccrualBox")
					err := w.repository.SendToAccrualBox(orders.orders, w.httpServer)
					if err != nil {
						fmt.Println("SendToAccrualBox err")
						w.l.Info("err in SendToAccrualBox():", err)
					}
				}
			}
		}()
	}
}

func (w *Pool) Stop() {
	w.once.Do(func() {
		close(w.shutDown)
		close(w.mainCh)
	})
	w.wg.Wait()
}
