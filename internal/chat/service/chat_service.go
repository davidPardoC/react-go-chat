package service

import (
	"fmt"
	"log"

	"github.com/davidPardoC/go-chat/cmd/chat/api/websocket/dtos"
	userRepo "github.com/davidPardoC/go-chat/internal/user/repository"
)

type ChatService struct {
	userRepository userRepo.IUserRepository
}

func NewChatService(userRepository userRepo.IUserRepository) *ChatService {
	return &ChatService{userRepository: userRepository}
}

func (s *ChatService) HandleTextMessage(chatEvent dtos.ChatEvent) {
	fmt.Printf("%v\n", chatEvent)
	receiver := Clients[uint(chatEvent.ReceiverId)]

	if receiver.Conn == nil {
		log.Printf("Not connection active for user with id %d \n", chatEvent.ReceiverId)
		return
	}

	var result = make(map[string]interface{})
	result["message"] = chatEvent.MessageText
	result["sender_id"] = 1234
	receiver.Conn.WriteJSON(result)
}
