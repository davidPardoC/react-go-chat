package main

import (
	"fmt"

	"github.com/davidPardoC/go-chat/internal/api/webscoket"
)

func main() {
	fmt.Println("Starting chat app")
	webscoket.StartWebSocketServer()
}
