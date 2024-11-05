package service

import "github.com/davidPardoC/go-chat/internal/chat/repository"

type ChatService struct {
	chatRepository *repository.IChatRepository
}

func NewChatService(chatRepository *repository.IChatRepository) *ChatService {
	return &ChatService{chatRepository: chatRepository}
}

func (s *ChatService) CreateNew() {}
