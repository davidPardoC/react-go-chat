package repository

import "github.com/davidPardoC/go-chat/internal/user/model"

type IUserRepository interface {
	GetUserByEmail(email string) (*model.User, error)
}
