package service

import (
	"fmt"

	"github.com/davidPardoC/go-chat/internal/chat/model"
	"github.com/davidPardoC/go-chat/internal/user/repository"
	"github.com/gorilla/websocket"
)

var Clients = make(map[int]model.Client)

type WebsocketService struct {
	userRepo repository.IUserRepository
}

func NewWebsocketService(userRepo repository.IUserRepository) *WebsocketService {
	return &WebsocketService{userRepo: userRepo}
}

func (s *WebsocketService) RegisterClient(userId uint, conn *websocket.Conn) {
	user, _ := s.userRepo.FindById(userId)
	newClient := model.Client{Conn: conn}
	Clients[int(user.ID)] = newClient
	fmt.Printf("Clients: %v", Clients)
}
