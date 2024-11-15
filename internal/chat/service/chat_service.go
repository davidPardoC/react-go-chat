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
	"github.com/davidPardoC/go-chat/pkg/cache"
	"github.com/davidPardoC/go-chat/pkg/constants"
	"gorm.io/gorm"
)

type ChatService struct {
	userRepository       userRepo.IUserRepository
	messageRepository    repository.IMessageRepository
	chatRepository       repository.IChatRepository
	chatMemberRepository repository.IChatMemberRepository
	cacheService         cache.ICacheService
}

func NewChatService(
	userRepository userRepo.IUserRepository,
	messageRepository repository.IMessageRepository,
	chatRepository repository.IChatRepository,
	chatMemberRepository repository.IChatMemberRepository,
	cacheService cache.ICacheService) *ChatService {

	return &ChatService{
		userRepository:       userRepository,
		messageRepository:    messageRepository,
		chatRepository:       chatRepository,
		chatMemberRepository: chatMemberRepository,
		cacheService:         cacheService,
	}
}

func (s *ChatService) CreateNew(chat model.Chat) (model.Chat, error) {
	chat, err := s.chatRepository.Create(chat)
	return chat, err
}

func (s *ChatService) HandleTextMessage(chatEvent dtos.ChatEvent, userID uint) {
	fmt.Printf("%v\n", chatEvent)
	receiver := Clients[chatEvent.RecipientID]

	chat, err := s.chatRepository.GetByChatMembers(int(userID), chatEvent.RecipientID)

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

	if receiver.Conn == nil {
		log.Printf("Not connection active for user with id %d \n", chatEvent.RecipientID)
		return
	}

	receiver.Conn.WriteJSON(outgoingMessage)
}

func (s *ChatService) CreateChatWithMembers(senderID int, recipientId int) (*model.Chat, error) {
	chat, err := s.CreateNew(model.Chat{IsSelf: senderID == recipientId})

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

	if senderID == recipientId {
		s.chatMemberRepository.Create(senderMember)
		return &chat, err
	}

	s.chatMemberRepository.Create(senderMember)
	s.chatMemberRepository.Create(recipientMember)

	return &chat, err
}

func (s *ChatService) GetUserChats(userId int) ([]model.ApiChat, error) {
	chats, err := s.chatRepository.FindByUserId(userId)
	return chats, err
}
