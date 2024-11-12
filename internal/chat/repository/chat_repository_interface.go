package repository

import "github.com/davidPardoC/go-chat/internal/chat/model"

type IChatRepository interface {
	Create() (model.Chat, error)
	FindById(chatId uint) (model.Chat, error)
	FindByUserId(userId int) ([]model.Chat, error)
}
