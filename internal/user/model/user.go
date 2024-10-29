package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint           `json:"id" gorm:"primaryKey`
	Username     string         `json:"username,omitempty"`
	Email        string         `json:"email,omitempty"`
	Password     string         `json:"password,omitempty"`
	RefreshToken sql.NullString `json:"refresh_token"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func (u *User) WithoutPassword() *User {
	u.Password = ""
	return u
}
