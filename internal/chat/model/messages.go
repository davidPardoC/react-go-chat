package model

import (
	userModel "github.com/davidPardoC/go-chat/internal/user/model"
	"gorm.io/gorm"
)

type Message struct {
	gorm.Model
	MessageText string         `json:"message_text"`
	Read        bool           `json:"read"`
	ChatID      uint           `json:"chat_id"`
	UserID      uint           `json:"user_id"`
	User        userModel.User `gorm:"foreignKey:UserID"`
}
