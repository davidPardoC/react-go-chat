package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ChatHandler struct{}

func NewChatHandler() *ChatHandler {
	return &ChatHandler{}
}

func (h *ChatHandler) CreateChat(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{})
}
