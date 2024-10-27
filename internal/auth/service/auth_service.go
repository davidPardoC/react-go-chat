package service

import (
	"github.com/davidPardoC/go-chat/pkg/user/model"
	"github.com/davidPardoC/go-chat/pkg/user/repository"
)

type AuthService struct {
	userRepo repository.IUserRepository
}

func NewAuthService(userRepo repository.IUserRepository) *AuthService {
	return &AuthService{}
}

func (s *AuthService) LoginUser(email string, password string) *model.User {
	user := s.userRepo.GetUserByEmail(email)
	return user
}
