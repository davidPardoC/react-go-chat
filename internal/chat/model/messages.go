package model

import (
	"time"

	userModel "github.com/davidPardoC/go-chat/internal/user/model"
)

type Message struct {
	ID          uint           `json:"id"`
	MessageText string         `json:"message_text"`
	Read        bool           `json:"read"`
	ChatID      uint           `json:"chat_id"`
	UserID      uint           `json:"user_id"`
	User        userModel.User `gorm:"foreignKey:UserID"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
}
