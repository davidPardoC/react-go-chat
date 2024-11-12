package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/davidPardoC/go-chat/cmd/chat/api/websocket/dtos"
	"github.com/davidPardoC/go-chat/internal/chat/model"
	"github.com/davidPardoC/go-chat/internal/chat/repository"
	userRepo "github.com/davidPardoC/go-chat/internal/user/repository"
	"github.com/davidPardoC/go-chat/pkg/constants"
	"gorm.io/gorm"
)

type ChatService struct {
	userRepository       userRepo.IUserRepository
	messageRepository    repository.IMessageRepository
	chatRepository       repository.IChatRepository
	chatMemberRepository repository.IChatMemberRepository
}

func NewChatService(
	userRepository userRepo.IUserRepository,
	messageRepository repository.IMessageRepository,
	chatRepository repository.IChatRepository,
	chatMemberRepository repository.IChatMemberRepository) *ChatService {

	return &ChatService{
		userRepository:       userRepository,
		messageRepository:    messageRepository,
		chatRepository:       chatRepository,
		chatMemberRepository: chatMemberRepository,
	}
}

func (s *ChatService) CreateNew() (model.Chat, error) {
	chat, err := s.chatRepository.Create()
	return chat, err
}

func (s *ChatService) HandleTextMessage(chatEvent dtos.ChatEvent, userID uint) {
	fmt.Printf("%v\n", chatEvent)
	receiver := Clients[chatEvent.RecipientID]

	if receiver.Conn == nil {
		log.Printf("Not connection active for user with id %d \n", chatEvent.RecipientID)
		return
	}

	chatMembersChat, err := s.chatMemberRepository.GetByChatMembers(int(userID), chatEvent.RecipientID)

	chat := model.Chat{ID: chatMembersChat.ID}

	if errors.Is(err, gorm.ErrRecordNotFound) {
		newChat, _ := s.CreateChatWithMembers(int(userID), chatEvent.RecipientID)
		chat = *newChat
	}

	message := model.Message{
		ChatID:      chat.ID,
		MessageText: chatEvent.MessageText,
		Read:        false,
		UserID:      userID,
	}

	savedMessage, err := s.messageRepository.Create(message)

	if err != nil {
		log.Printf("Cannot save and send message %v due to %v\n", savedMessage, err)
	}

	outgoingChatEvent := dtos.ChatEvent{
		MessageText: savedMessage.MessageText,
		RecipientID: chatEvent.RecipientID,
		ChatId:      int(chat.ID),
	}

	data, err := json.Marshal(outgoingChatEvent)
	if err != nil {
		log.Printf("Cannot marshal outgoing chat event: %v\n", err)
		return
	}

	outgoingMessage := dtos.MessageDTO{
		Event:  constants.TEXT_MESSAGE,
		UserID: userID,
		Data:   json.RawMessage(data),
	}

	receiver.Conn.WriteJSON(outgoingMessage)
}

func (s *ChatService) CreateChatWithMembers(senderID int, recipientId int) (*model.Chat, error) {
	chat, err := s.CreateNew()

	if err != nil {
		return nil, err
	}

	sender, err := s.userRepository.FindById(uint(senderID))

	if err != nil {
		return nil, err
	}

	recipient, err := s.userRepository.FindById(uint(recipientId))

	if err != nil {
		return nil, err
	}

	senderMember := model.ChatMember{Chat: chat, UserID: sender.ID}
	recipientMember := model.ChatMember{Chat: chat, UserID: recipient.ID}

	s.chatMemberRepository.Create(senderMember)
	s.chatMemberRepository.Create(recipientMember)

	return &chat, err
}

func (s *ChatService) GetUserChats(userId int) ([]model.Chat, error) {
	chats, err := s.chatRepository.FindByUserId(userId)
	return chats, err
}
