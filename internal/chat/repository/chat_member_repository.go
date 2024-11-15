package repository

import (
	"github.com/davidPardoC/go-chat/internal/chat/model"
	"gorm.io/gorm"
)

type ChatMemberRepository struct {
	db *gorm.DB
}

func NewChatMemberRepository(db *gorm.DB) *ChatMemberRepository {
	return &ChatMemberRepository{db: db}
}

func (r *ChatMemberRepository) Create(member model.ChatMember) (model.ChatMember, error) {
	result := r.db.Create(&member)
	return member, result.Error
}
