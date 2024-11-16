package dtos

type GetChatsUri struct {
	UserId int `uri:"user_id"`
}

type GetFullChatUri struct {
	ID int `uri:"id"`
}
