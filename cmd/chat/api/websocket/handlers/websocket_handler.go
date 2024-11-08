package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/davidPardoC/go-chat/cmd/chat/api/websocket/dtos"
	"github.com/davidPardoC/go-chat/internal/chat/service"
	"github.com/davidPardoC/go-chat/pkg/constants"
	"github.com/davidPardoC/go-chat/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

var BroadCast = make(chan dtos.MessageDTO)

type WsHandler struct {
	websockerService *service.WebsocketService
	chatService      *service.ChatService
}

func NewWsHandler(websockerService *service.WebsocketService, chatService *service.ChatService) *WsHandler {
	return &WsHandler{websockerService: websockerService, chatService: chatService}
}

func checkOrigin(r *http.Request) bool {
	values := r.URL.Query()
	token := values.Get("token")

	isValid, _ := utils.IsTokenValid(token)

	return isValid
}

func getTokenClaims(r *http.Request) jwt.MapClaims {
	values := r.URL.Query()
	token := values.Get("token")

	_, claims := utils.IsTokenValid(token)
	return claims
}

var upgrader = websocket.Upgrader{
	CheckOrigin: checkOrigin,
}

func (h *WsHandler) Handle(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Print("upgrade error:", err)
		return
	}
	defer conn.Close()

	claims := getTokenClaims(ctx.Request)
	userId, _ := claims.GetSubject()

	userIdUint, err := strconv.ParseUint(userId, 10, 32)

	if err != nil {
		log.Printf("error converting userId to uint: %v", err)
		return
	}

	h.websockerService.RegisterClient(uint(userIdUint), conn)

	for {
		var msg dtos.MessageDTO
		err := conn.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		BroadCast <- msg
	}
}

func (h *WsHandler) HandleMessages() {
	for {
		msg := <-BroadCast

		switch msg.Event {
		case constants.TEXT_MESSAGE:

			var chatEvent dtos.ChatEvent
			err := json.Unmarshal(msg.Data, &chatEvent)

			if err != nil {
				log.Printf("Error unmarshaling %v\n", msg)
			}
			h.chatService.HandleTextMessage(chatEvent)

		default:
			fmt.Println("Received unknown event type")
		}
	}
}
