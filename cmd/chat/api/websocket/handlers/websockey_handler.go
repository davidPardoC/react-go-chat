package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davidPardoC/go-chat/cmd/chat/api/websocket/dtos"
	"github.com/davidPardoC/go-chat/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

var BroadCast = make(chan dtos.MessageDTO[interface{}])

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

func WsHandler(ctx *gin.Context) {
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Print("upgrade error:", err)
		return
	}
	defer c.Close()

	claims := getTokenClaims(ctx.Request)

	userId, _ := claims.GetSubject()

	fmt.Println("claims: %v, userId: %s", claims, userId)

	for {
		var msg dtos.MessageDTO[interface{}]
		err := c.ReadJSON(&msg)
		if err != nil {
			fmt.Println(err)
			return
		}
		BroadCast <- msg
	}
}

func HandleMessages() {
	for {
		msg := <-BroadCast
		switch msg.Event {
		case "message":
			fmt.Printf("Some chat message, data: %v \n", msg.Data)
			if data, ok := msg.Data.(dtos.ChatEvent); ok {
				fmt.Printf("Chat message received, data: %v \n", data.MessageText)
			}
		default:
			fmt.Println("Received unknown event type")
		}
	}
}
