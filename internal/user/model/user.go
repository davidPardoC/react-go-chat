package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint           `json:"id" gorm:"primaryKey"`
	Username     string         `json:"username,omitempty" gorm:"unique"`
	Email        string         `json:"email,omitempty" gorm:"unique"`
	Password     string         `json:"password,omitempty"`
	RefreshToken sql.NullString `json:"refresh_token"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
}

func (u *User) WithoutPassword() *User {
	u.Password = ""
	return u
}

func (u *User) RemoveRefreshToken() *User {
	u.RefreshToken = sql.NullString{}
	return u
}
