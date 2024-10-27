package handlers

import (
	"github.com/davidPardoC/go-chat/internal/auth/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authServcices service.AuthService
}

func NewAuthHandler(authServcices service.AuthService) AuthHandler {
	return AuthHandler{authServcices: authServcices}
}

func (h *AuthHandler) SignupHandler(c *gin.Context) {}

func (h *AuthHandler) Loginhandler(c *gin.Context) {}
