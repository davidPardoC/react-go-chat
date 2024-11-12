package model

import (
	"time"
)

type Message struct {
	ID          uint      `json:"id"`
	MessageText string    `json:"message_text"`
	Read        bool      `json:"read"`
	ChatID      uint      `json:"chat_id"`
	UserID      uint      `json:"user_id"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
