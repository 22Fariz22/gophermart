package entity

import "time"

type Order struct {
	ID             uint32
	UserID         uint32
	Number         uint32
	UploadedAtDate time.Time
}
