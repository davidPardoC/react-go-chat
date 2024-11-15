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

func (app *ServerApp) StartHttpServer() {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	config.AllowMethods = []string{"HEAD", "OPTIONS", "GET", "POST", "PUT", "DELETE", "PATCH"}
	config.AllowHeaders = []string{"Origin", "Authorization", "Content-Type", "Referer", "User-Agent"}
	config.AllowCredentials = true
	r.Use(cors.New(config))

	router.SetHealthRouter(r)
	router.SetAuthRouter(r, app.db)
	router.SetUsersRouter(r, app.db)
	router.SetChatRouter(r, app.db)

	websocket.StartWebSocketServer(r, app.db)

	r.Run(":5500")
}
