package main

import (
	"fmt"

	"github.com/davidPardoC/go-chat/cmd/chat/handlers/api"
	"github.com/davidPardoC/go-chat/cmd/chat/handlers/websocket"
)

func main() {
	fmt.Println("Starting chat app")
	websocket.StartWebSocketServer()
	api.StartHttpServer()
}
