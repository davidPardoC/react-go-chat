package repository

import "github.com/davidPardoC/go-chat/internal/chat/model"

type IChatMemberRepository interface {
	Create(model.ChatMember) (model.ChatMember, error)
	GetByChatMembers(user1 int, user2 int) (model.ChatMember, error)
}
