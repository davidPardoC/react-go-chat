package api

import (
	"github.com/davidPardoC/go-chat/cmd/chat/handlers/api/router"
	"github.com/davidPardoC/go-chat/cmd/chat/handlers/websocket"
	"github.com/gin-gonic/gin"
)

func StartHttpServer() {
	r := gin.Default()
	router.SetHealthRouter(r)
	websocket.StartWebSocketServer(r)
	r.Run(":5000")
}
