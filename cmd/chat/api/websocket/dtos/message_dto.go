package dtos

import "encoding/json"

type MessageDTO struct {
	Event  string          `json:"event"`
	UserID uint            `json:"user_id"`
	Data   json.RawMessage `json:"data"`
}

type ChatEvent struct {
	MessageText string `json:"message_text"`
	ReceiverId  int    `json:"receiver_id"`
	ChatId      int    `json:"chat_id"`
}
