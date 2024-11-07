package service

import (
	"github.com/davidPardoC/go-chat/internal/user/model"
	"github.com/davidPardoC/go-chat/internal/user/repository"
)

type UserService struct {
	userRepository repository.IUserRepository
}

func NewUserService(userRepository repository.IUserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (s *UserService) GetAll() ([]model.User, error) {
	users, err := s.userRepository.FindAll()
	mappedUsers := []model.User{}
	for _, user := range users {
		mappedUsers = append(mappedUsers, *user.WithoutPassword().RemoveRefreshToken())
	}
	return mappedUsers, err
}
