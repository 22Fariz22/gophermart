package entity

import "time"

type Order struct {
	ID     uint32
	UserID uint32
	Number uint32
	//Accrual    uint32
	Status     string
	UploadedAt time.Time
}
