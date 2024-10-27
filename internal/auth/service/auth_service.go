package service

import (
	"github.com/davidPardoC/go-chat/internal/user/model"
	"github.com/davidPardoC/go-chat/internal/user/repository"
)

type AuthService struct {
	userRepo repository.IUserRepository
}

func NewAuthService(userRepo repository.IUserRepository) AuthService {
	return AuthService{}
}

func (s *AuthService) LoginUser(email string, password string) *model.User {
	user, _ := s.userRepo.GetUserByEmail(email)
	return user
}
