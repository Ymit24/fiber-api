package models

import "time"

type User struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Email       string `json:"email"`
	Username    string `json:"username"`
	PasswordEnc string `json:"password_enc"`
	CreatedAt   time.Time
}
