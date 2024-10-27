package repository

import "github.com/davidPardoC/go-chat/pkg/user/model"

type IUserRepository interface {
	GetUserByEmail(email string) *model.User
}
