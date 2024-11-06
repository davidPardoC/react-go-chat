package api

import (
	"github.com/davidPardoC/go-chat/cmd/chat/api/http/router"
	"github.com/davidPardoC/go-chat/cmd/chat/api/websocket"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type ServerApp struct {
	db *gorm.DB
}

func NewServerApp(db *gorm.DB) *ServerApp {
	return &ServerApp{db: db}
}

func (a *ServerApp) StartHttpServer() {
	r := gin.Default()

	r.Use(cors.Default())

	router.SetHealthRouter(r)
	router.SetAuthRouter(r, a.db)
	router.SetUsersRouter(r, a.db)

	websocket.StartWebSocketServer(r)

	r.Run(":5500")
}
