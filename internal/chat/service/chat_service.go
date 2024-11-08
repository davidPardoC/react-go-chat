package service

import (
	"fmt"
	"log"

	"github.com/davidPardoC/go-chat/cmd/chat/api/websocket/dtos"
	"github.com/davidPardoC/go-chat/internal/chat/model"
	"github.com/davidPardoC/go-chat/internal/chat/repository"
	userRepo "github.com/davidPardoC/go-chat/internal/user/repository"
)

type ChatService struct {
	userRepository    userRepo.IUserRepository
	messageRepository repository.IMessageRepository
}

func NewChatService(userRepository userRepo.IUserRepository, messageRepository repository.IMessageRepository) *ChatService {
	return &ChatService{
		userRepository:    userRepository,
		messageRepository: messageRepository,
	}
}

func (s *ChatService) HandleTextMessage(chatEvent dtos.ChatEvent, userID uint) {
	fmt.Printf("%v\n", chatEvent)
	receiver := Clients[uint(chatEvent.ReceiverId)]

	if receiver.Conn == nil {
		log.Printf("Not connection active for user with id %d \n", chatEvent.ReceiverId)
		return
	}

	message := model.Message{ChatID: uint(chatEvent.ChatId), MessageText: chatEvent.MessageText, Read: false, UserID: userID}

	savedMessage, err := s.messageRepository.Create(message)

	if err != nil {
		log.Printf("Cannot save and send message %v due to %v\n", savedMessage, err)
	}

	receiver.Conn.WriteJSON(savedMessage)
}
