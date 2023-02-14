package entity

import "time"

type History struct {
	ID          string
	UserID      string
	Number      string
	Sum         int
	ProcessedAt time.Time
}
