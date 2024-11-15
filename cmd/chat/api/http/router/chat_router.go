package router

import (
	"github.com/davidPardoC/go-chat/cmd/chat/api/http/handlers"
	"github.com/davidPardoC/go-chat/cmd/chat/api/http/middlewares"
	"github.com/davidPardoC/go-chat/internal/chat/repository"
	"github.com/davidPardoC/go-chat/internal/chat/service"
	userRepo "github.com/davidPardoC/go-chat/internal/user/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetChatRouter(r *gin.Engine, db *gorm.DB) {

	userRepo := userRepo.NewUserRepository(db)
	messagesRepo := repository.NewMessageRepository(db)
	chatRepo := repository.NewChatRepository(db)
	chatMemberRepo := repository.NewChatMemberRepository(db)

	chatService := service.NewChatService(userRepo, messagesRepo, chatRepo, chatMemberRepo)
	handler := handlers.NewChatHandler(chatService)

	usersV1 := r.Group("/v1/chats")
	{
		usersV1.Use(middlewares.AuthMiddleware()).GET("/", handler.GetChatList)
	}
}
