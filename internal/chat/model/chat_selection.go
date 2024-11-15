package model

type ApiUser struct {
	ID       uint   `json:"id,omitempty"`
	Username string `json:"username,omitempty"`
}

func (ApiUser) TableName() string {
	return "users"
}

type ApiChatMember struct {
	ID     uint    `json:"id"`
	ChatID int     `json:"chat_id"`
	UserID uint    `json:"user_id"`
	User   ApiUser `json:"user"`
}

func (ApiChatMember) TableName() string {
	return "chat_members"
}

type ApiMessage struct {
	ID          uint   `json:"id,omitempty"`
	MessageText string `json:"message_text,omitempty"`
	Read        bool   `json:"read,omitempty"`
	ChatID      uint   `json:"chat_id,omitempty"`
}

func (ApiMessage) TableName() string {
	return "messages"
}

type ApiChat struct {
	ID          uint            `json:"id"`
	Messages    []ApiMessage    `json:"messages" gorm:"foreignKey:ChatID"`
	ChatMembers []ApiChatMember `json:"chat_members" gorm:"foreignKey:ChatID"`
}
