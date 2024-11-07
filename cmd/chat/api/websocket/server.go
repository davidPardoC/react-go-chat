package websocket

import (
	"github.com/davidPardoC/go-chat/cmd/chat/api/websocket/handlers"
	"github.com/gin-gonic/gin"
)

func StartWebSocketServer(r *gin.Engine) {
	go handlers.HandleMessages()
	r.GET("/ws", func(c *gin.Context) {
		handlers.WsHandler(c)
	})
}
