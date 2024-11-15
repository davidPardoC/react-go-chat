package model

import (
	"time"

	"github.com/davidPardoC/go-chat/internal/user/model"
)

type Message struct {
	ID          uint       `json:"id,omitempty"`
	MessageText string     `json:"message_text,omitempty"`
	Read        bool       `json:"read,omitempty"`
	ChatID      uint       `json:"chat_id,omitempty"`
	UserID      uint       `json:"user_id,omitempty"`
	User        model.User `json:"user,omitempty"`
	CreatedAt   time.Time  `json:"created_at,omitempty"`
	UpdatedAt   time.Time  `json:"updated_at,omitempty"`
}
