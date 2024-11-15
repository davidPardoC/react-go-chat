package router

import (
	"github.com/davidPardoC/go-chat/cmd/chat/api/http/handlers"
	"github.com/davidPardoC/go-chat/cmd/chat/api/http/middlewares"
	"github.com/davidPardoC/go-chat/internal/user/repository"
	"github.com/davidPardoC/go-chat/internal/user/service"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetUsersRouter(r *gin.Engine, db *gorm.DB) {

	userRepos := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepos)
	handler := handlers.NewUsersHandler(userService)

	usersV1 := r.Group("/v1/users")
	{
		usersV1.Use(middlewares.AuthMiddleware()).GET("/", handler.GetAll)
	}
}
