package dtos

import "encoding/json"

type MessageDTO struct {
	Event  string          `json:"event"`
	UserID uint            `json:"user_id"`
	Data   json.RawMessage `json:"data"`
}

type ChatEvent struct {
	MessageText string `json:"message_text"`
	RecipientID int    `json:"recipient_id"`
	ChatId      int    `json:"chat_id"`
}
