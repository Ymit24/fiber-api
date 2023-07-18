package models

import "time"

type TimeSheet struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	Title     string `json:"title"`
	CreatedAt time.Time
}
