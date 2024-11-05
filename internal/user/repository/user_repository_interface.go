package repository

import (
	"github.com/davidPardoC/go-chat/internal/user/model"
)

type IUserRepository interface {
	GetUserByEmail(email string) (*model.User, error)
	CreateUser(user *model.User) (*model.User, error)
	UpdateRefresToken(user model.User, refreshToken string) (model.User, error)
	FindAll() ([]model.User, error)
}
