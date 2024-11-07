package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davidPardoC/go-chat/internal/chat/model"
	"github.com/davidPardoC/go-chat/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gorilla/websocket"
)

var broadcast = make(chan model.Message)

func checkOrigin(r *http.Request) bool {
	values := r.URL.Query()
	token := values.Get("token")

	isValid, _ := utils.IsTokenValid(token)

	return isValid
}

func getTokenClaims(r *http.Request) *jwt.MapClaims {
	values := r.URL.Query()
	token := values.Get("token")

	_, claims := utils.IsTokenValid(token)

	return claims
}

func handleMessages() {
	for {
		msg := <-broadcast
		fmt.Printf("%v", msg)
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: checkOrigin,
}

func wsHandler(ctx *gin.Context) {
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Print("upgrade error:", err)
		return
	}
	defer c.Close()

	claims := getTokenClaims(ctx.Request)

	fmt.Printf("%v", claims)

}

func StartWebSocketServer(r *gin.Engine) {
	go handleMessages()
	r.GET("/ws", func(c *gin.Context) {
		wsHandler(c)
	})
}
