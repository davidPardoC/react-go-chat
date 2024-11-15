package model

import (
	"time"

	"github.com/davidPardoC/go-chat/internal/user/model"
)

type ChatMember struct {
	ID        uint       `json:"id"`
	ChatID    int        `json:"chat_id"`
	Chat      Chat       `json:"chat,omitempty"`
	UserID    uint       `json:"user_id"`
	User      model.User `json:"user"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
