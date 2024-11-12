package model

import "time"

type Chat struct {
	ID        uint      `json:"id" gorm:"primaryKey,index"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Messages  []Message `json:"messages"`
}
