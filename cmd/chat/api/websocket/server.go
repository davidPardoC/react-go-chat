package websocket

import (
	"github.com/davidPardoC/go-chat/cmd/chat/api/websocket/handlers"
	"github.com/davidPardoC/go-chat/internal/chat/repository"
	"github.com/davidPardoC/go-chat/internal/chat/service"
	"github.com/davidPardoC/go-chat/internal/config"
	userRepo "github.com/davidPardoC/go-chat/internal/user/repository"
	"github.com/davidPardoC/go-chat/pkg/cache"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func StartWebSocketServer(r *gin.Engine, db *gorm.DB) {

	userRepo := userRepo.NewUserRepository(db)
	messagesRepo := repository.NewMessageRepository(db)
	chatRepo := repository.NewChatRepository(db)
	chatMemberRepo := repository.NewChatMemberRepository(db)

	redisClient := config.NewRedisClient()
	cacheService := cache.NewRedisCacheService(redisClient.Rdb)

	chatService := service.NewChatService(userRepo, messagesRepo, chatRepo, chatMemberRepo, cacheService)
	websocketService := service.NewWebsocketService(userRepo)

	handler := handlers.NewWsHandler(websocketService, chatService)

	go handler.HandleMessages()
	r.GET("/ws", handler.Handle)
}
