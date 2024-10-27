package main

import (
	"github.com/davidPardoC/go-chat/cmd/chat/handlers/api"
	"github.com/davidPardoC/go-chat/cmd/chat/handlers/websocket"
)

func main() {
	go websocket.StartWebSocketServer()
	api.StartHttpServer()

}
