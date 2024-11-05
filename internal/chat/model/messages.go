package model

import "gorm.io/gorm"

type Messages struct {
	gorm.Model
	MessageText string `json:"message_text"`
	Read        bool   `json:"read"`
	ChatID      uint
}
