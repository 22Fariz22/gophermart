package entity

import "time"

type Order struct {
	ID     string
	UserID string
	Number uint32
	//Accrual    uint32
	Status     string
	UploadedAt time.Time
}
