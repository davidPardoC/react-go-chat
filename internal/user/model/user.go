package model

import (
	"database/sql"
	"time"
)

type User struct {
	ID           uint            `json:"id,omitempty" gorm:"primaryKey"`
	Username     string          `json:"username,omitempty" gorm:"unique"`
	Email        string          `json:"email,omitempty" gorm:"unique"`
	Password     string          `json:"password,omitempty"`
	RefreshToken *sql.NullString `json:"refresh_token,omitempty"`
	CreatedAt    time.Time       `json:"created_at,omitempty"`
	UpdatedAt    time.Time       `json:"updated_at,omitempty"`
}

func (u *User) WithoutPassword() *User {
	u.Password = ""
	return u
}

func (u *User) RemoveRefreshToken() *User {
	u.RefreshToken = nil
	return u
}
