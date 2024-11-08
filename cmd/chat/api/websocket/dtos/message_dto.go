package dtos

import "encoding/json"

type MessageDTO struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

type ChatEvent struct {
	MessageText string `json:"message_text"`
	ReceiverId  int    `json:"receiver_id"`
}
