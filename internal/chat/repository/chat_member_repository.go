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

func (r *ChatMemberRepository) GetByChatMembers(user1 int, user2 int) (model.ChatMember, error) {
	var chatMember model.ChatMember

	result :=
		r.db.Table("chat_members").
			Joins("JOIN chat_members AS cm1 ON cm1.chat_id = chat_members.chat_id").
			Joins("JOIN chat_members AS cm2 ON cm2.chat_id = chat_members.chat_id").
			Where("cm1.id  = ?", user1).Where("cm2.id  = ?", user2).First(&chatMember)

	return chatMember, result.Error
}
