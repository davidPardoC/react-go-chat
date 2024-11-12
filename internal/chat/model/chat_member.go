package model

import (
	"time"
)

type ChatMember struct {
	ID        uint `json:"id"`
	ChatID    int
	Chat      Chat
	UserID    uint
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
