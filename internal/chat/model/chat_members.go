package model

import "gorm.io/gorm"

type ChatMembers struct {
	gorm.Model
	ChatID uint
	UserID uint
}
