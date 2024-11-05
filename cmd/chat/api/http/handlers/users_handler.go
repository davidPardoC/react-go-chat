package handlers

import (
	"net/http"

	"github.com/davidPardoC/go-chat/internal/user/service"
	"github.com/gin-gonic/gin"
)

type UsersHandler struct {
	userService *service.UserService
}

func NewUsersHandler(userService *service.UserService) *UsersHandler {
	return &UsersHandler{userService: userService}
}

func (h *UsersHandler) GetAll(c *gin.Context) {
	users, _ := h.userService.GetAll()
	c.JSON(http.StatusOK, users)
}