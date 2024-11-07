package dtos

type MessageDTO[D any] struct {
	Event string `json:"event"`
	Data  D
}

type ChatEvent struct {
	MessageText string `json:"message_text"`
}
