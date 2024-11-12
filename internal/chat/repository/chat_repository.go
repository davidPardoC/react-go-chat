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
	chat := model.Chat{}

	result := r.db.Where("id = ? ", chatId).First(&chat)

	return chat, result.Error
}

func (r *ChatRepository) FindByUserId(userId int) ([]model.Chat, error) {
	chats := []model.Chat{}
	r.db.Model(&model.Chat{}).Preload("Messages", func(tx *gorm.DB) *gorm.DB {
		return tx.Order("created_at DESC").Limit(1)
	}).Joins("LEFT join chat_members on chat_members.user_id = ?", userId).Find(&chats)
	return chats, nil
}
