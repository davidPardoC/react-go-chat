package router

import (
	"github.com/davidPardoC/go-chat/cmd/chat/api/http/handlers"
	"github.com/davidPardoC/go-chat/internal/auth/service"
	"github.com/davidPardoC/go-chat/internal/user/repository"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetAuthRouter(r *gin.Engine, db *gorm.DB) {

	userRepo := repository.NewUserRepository(db)
	authService := service.NewAuthService(userRepo)
	handler := handlers.NewAuthHandler(authService)

	authV1 := r.Group("/v1/auth")
	{
		authV1.POST("/signup", handler.SignupHandler)
		authV1.POST("/login", handler.Loginhandler)
	}
}
