package handlers

import (
	"net/http"
	"strconv"

	"github.com/davidPardoC/go-chat/cmd/chat/api/http/dtos"
	"github.com/davidPardoC/go-chat/internal/chat/service"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

type ChatHandler struct {
	chatService *service.ChatService
}

func NewChatHandler(chatService *service.ChatService) *ChatHandler {
	return &ChatHandler{chatService: chatService}
}

func (h *ChatHandler) CreateChat(ctx *gin.Context) {
	ctx.JSON(http.StatusCreated, gin.H{})
}

func (h *ChatHandler) GetChatList(ctx *gin.Context) {
	user := ctx.MustGet("user").(jwt.MapClaims)

	var queryDto dtos.GetChatsUri

	if err := ctx.ShouldBindUri(&queryDto); err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	userId, _ := user.GetSubject()

	userIdUint, _ := strconv.ParseUint(userId, 10, 32)

	chats, _ := h.chatService.GetUserChats(int(userIdUint))

	ctx.JSON(http.StatusOK, chats)
}
