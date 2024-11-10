package model

import (
	"time"

	"github.com/davidPardoC/go-chat/internal/user/model"
)

type ChatMember struct {
	ID        uint `json:"id"`
	ChatID    int
	Chat      Chat
	UserID    int
	User      model.User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
