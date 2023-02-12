package entity

import "time"

type Balance struct {
	ID           string
	OrderID      string
	UserID       string
	Accrual      uint32
	Status       uint8
	UploadedAt   time.Time
	WithdrawDate time.Time
}
