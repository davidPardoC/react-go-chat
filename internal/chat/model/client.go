package model

import (
	"github.com/davidPardoC/go-chat/internal/user/model"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn *websocket.Conn
	User *model.User
}
