package handlers

import (
	"net/http"

	"github.com/davidPardoC/go-chat/cmd/chat/api/http/dtos"
	"github.com/davidPardoC/go-chat/internal/auth/service"
	"github.com/gin-gonic/gin"
)

type AuthHandler struct {
	authServcices *service.AuthService
}

func NewAuthHandler(authServcices *service.AuthService) *AuthHandler {
	return &AuthHandler{authServcices: authServcices}
}

func (h *AuthHandler) SignupHandler(c *gin.Context) {
	var signUp dtos.SignUpDto
	if err := c.ShouldBindJSON(&signUp); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}
	user, err := h.authServcices.SignupUser(signUp)

	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *AuthHandler) Loginhandler(c *gin.Context) {
	var login dtos.LoginDto

	if err := c.ShouldBindJSON(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	credentials, err := h.authServcices.LoginUser(login.Email, login.Password)

	if err != nil {
		c.JSON(err.StatusCode, gin.H{"error": err.Message})
		return
	}

	c.JSON(http.StatusOK, credentials)
}
