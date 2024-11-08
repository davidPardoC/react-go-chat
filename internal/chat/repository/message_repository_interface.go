package repository

import "github.com/davidPardoC/go-chat/internal/chat/model"

type IMessageRepository interface {
	GetAllByChatId(chatId uint) ([]model.Message, error)
	Create(message model.Message) (model.Message, error)
}
