package repository

import "github.com/davidPardoC/go-chat/internal/chat/model"

type IChatRepository interface {
	Create(chat model.Chat) (model.Chat, error)
	FindById(chatId uint) (model.Chat, error)
	FindByUserId(userId int) ([]model.ApiChat, error)
	GetByChatMembers(user1 int, user2 int) (model.Chat, error)
}
