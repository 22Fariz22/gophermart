package worker

import (
	"errors"
	"github.com/22Fariz22/gophermart/internal/entity"
	"github.com/22Fariz22/gophermart/pkg/logger"
	"sync"
	"time"
)

type Pool struct {
	wg         sync.WaitGroup
	once       sync.Once
	shutDown   chan struct{}
	mainCh     chan workerData
	repository UseCase
	l          logger.Interface
}

func NewWorkerPool(repo UseCase, l logger.Interface) *Pool {
	return &Pool{
		wg:         sync.WaitGroup{},
		once:       sync.Once{},
		shutDown:   make(chan struct{}),
		mainCh:     make(chan workerData, 10),
		repository: repo,
		l:          l,
	}
}

// создаем функцию которая каждые 2 мин забирает из таблицы ордеры со статусом NEW и кдадет их в каналы
func CollectNewOrders(uc UseCase, l logger.Interface) []*entity.Order {
	workers := NewWorkerPool(uc, l)

	for {
		time.Sleep(120 * time.Second)

		workers.RunWorkers(100)
		newOrders, err := workers.repository.CheckNewOrders()
		if err != nil {
			l.Info("err in CheckNewOrders(): ", err)
		}
		workers.AddJob(newOrders)
	}

	return nil
}

type workerData struct {
	orders []*entity.Order
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
	for i := 0; i < count; i++ {
		w.wg.Add(1)
		go func() {
			defer w.wg.Done()
			for {
				select {
				case <-w.shutDown:
					//w.l.Info("channels are shutdown.")
					return
				case orders, ok := <-w.mainCh:
					if !ok {
						return
					}
					err := w.repository.SendToOrdersCannels(orders.orders)
					if err != nil {
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
