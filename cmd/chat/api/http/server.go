package api

import (
	"github.com/davidPardoC/go-chat/cmd/chat/api/http/router"
	"github.com/davidPardoC/go-chat/cmd/chat/api/websocket"
	"github.com/gin-gonic/gin"
)

func StartHttpServer() {
	r := gin.Default()

	router.SetHealthRouter(r)
	router.SetAuthRouter(r)

	websocket.StartWebSocketServer(r)

	r.Run(":5000")
}
