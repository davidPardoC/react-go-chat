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

func (r *ChatRepository) Create(chat model.Chat) (model.Chat, error) {

	result := r.db.Create(&chat)

	return chat, result.Error
}

func (r *ChatRepository) FindById(chatId uint) (model.Chat, error) {
	chat := model.Chat{}

	result := r.db.Model(&model.Chat{}).Preload("Messages").Where("id = ? ", chatId).First(&chat)

	return chat, result.Error
}

func (r *ChatRepository) FindByUserId(userId int) ([]model.ApiChat, error) {
	chats := []model.ApiChat{}

	subQuery := r.db.Table("messages").
		Select("MAX(id) as id").
		Group("chat_id")

	r.db.
		Preload("Messages", func(tx *gorm.DB) *gorm.DB {
			return tx.Where("id IN (?)", subQuery)
		}).
		Preload("ChatMembers.User", func(tx *gorm.DB) *gorm.DB {
			return tx.Select("id", "username")
		}).
		Model(&model.Chat{}).
		Joins("LEFT join chat_members ON chat_members.chat_id = chats.id").
		Where("user_id = ?", userId).Find(&chats)

	return chats, nil
}

func (r *ChatRepository) GetByChatMembers(user1 int, user2 int) (model.Chat, error) {
	var chat model.Chat

	if user1 == user2 {
		result := r.db.Table("chats").
			Joins("JOIN chat_members ON chats.id = chat_members.chat_id").
			Where("chat_members.user_id = ?", user1).Where("chats.is_self = ?", true).First(&chat)
		return chat, result.Error
	}

	subquery := r.db.Table("chat_members").
		Select("chat_id").
		Where("user_id = ?", user1).
		Where("chat_id IN (?)",
			r.db.Table("chat_members").Select("chat_id").Where("user_id = ?", user2),
		)

	result := r.db.Table("chats").Where("id IN (?)", subquery).First(&chat)

	return chat, result.Error
}
