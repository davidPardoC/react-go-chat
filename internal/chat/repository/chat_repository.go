package repository

import (
	"github.com/davidPardoC/go-chat/internal/chat/model"
	"gorm.io/gorm"
)

type ChatRepository struct {
	db *gorm.DB
}

func NewChatRepository(db *gorm.DB) *ChatRepository {
	return &ChatRepository{db: db}
}

func (r *ChatRepository) Create() (model.Chat, error) {
	chat := model.Chat{}

	result := r.db.Create(&chat)

	return chat, result.Error
}

func (r *ChatRepository) FindById(chatId uint) (model.Chat, error) {
	chat := model.Chat{Model: gorm.Model{ID: chatId}}

	result := r.db.First(&chat)

	return chat, result.Error
}
