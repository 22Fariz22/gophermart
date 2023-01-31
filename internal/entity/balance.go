package entity

import "time"

type Balance struct {
	ID             uint32
	UserID         uint32
	OrderID        uint32
	Accrual        uint32
	Status         uint8
	UploadedAtDate time.Time
	WithdrawDate   time.Time
}
