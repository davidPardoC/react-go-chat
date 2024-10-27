package repository

import "github.com/davidPardoC/go-chat/internal/user/model"

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	return nil, nil
}
