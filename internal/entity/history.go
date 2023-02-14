package entity

import "time"

type History struct {
	ID           string
	OrderID      string
	UserID       string
	Accrual      uint32
	Status       uint8
	Processed_At time.Time
}
