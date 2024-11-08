package repository

import (
	"github.com/davidPardoC/go-chat/internal/user/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) GetUserByEmail(email string) (*model.User, error) {
	user := model.User{}

	result := r.db.Where("email = ?", email).First(&user)

	return &user, result.Error
}

func (r *UserRepository) CreateUser(user *model.User) (*model.User, error) {
	result := r.db.Create(&user)

	return user, result.Error
}

func (r *UserRepository) UpdateRefresToken(user model.User, refreshToken string) (model.User, error) {
	result := r.db.Model(&user).Update("refresh_token", refreshToken)
	return user, result.Error
}

func (r *UserRepository) FindAll() ([]model.User, error) {
	users := []model.User{}
	result := r.db.Find(&users)
	return users, result.Error
}

func (r *UserRepository) FindById(ID uint) (model.User, error) {
	user := model.User{ID: ID}
	result := r.db.First(&user)
	return user, result.Error
}
