package repository

import "github.com/davidPardoC/go-chat/internal/chat/model"

type IChatMemberRepository interface {
	Create(model.ChatMember) (model.ChatMember, error)
}
