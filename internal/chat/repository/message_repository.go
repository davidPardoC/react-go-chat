package repository

import (
	"github.com/davidPardoC/go-chat/internal/chat/model"
	"gorm.io/gorm"
)

type MessageRepository struct {
	db *gorm.DB
}

func NewMessageRepository(db *gorm.DB) *MessageRepository {
	return &MessageRepository{db: db}
}

func (r *MessageRepository) GetAllByChatId(chatId uint) ([]model.Message, error) {
	messages := []model.Message{}

	result := r.db.Where("chat_id = ?", chatId).Order("created_at DESC").Find(&messages)

	return messages, result.Error
}

func (r *MessageRepository) Create(message model.Message) (model.Message, error) {
	result := r.db.Create(&message)
	return message, result.Error
}
