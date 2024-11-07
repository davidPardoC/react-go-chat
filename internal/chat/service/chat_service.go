package service

import userRepo "github.com/davidPardoC/go-chat/internal/user/repository"

type ChatService struct {
	userRepository *userRepo.IUserRepository
}

func NewChatService(userRepository *userRepo.IUserRepository) *ChatService {
	return &ChatService{}
}
