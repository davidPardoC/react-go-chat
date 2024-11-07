package model

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	MessageText string `json:"message_text"`
	Read        bool   `json:"read"`
	ChatID      uint
}
