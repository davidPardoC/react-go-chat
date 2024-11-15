package repository

import "github.com/davidPardoC/go-chat/internal/chat/model"

type IChatMemberRepository interface {
	Create(member model.ChatMember) (model.ChatMember, error)
}
