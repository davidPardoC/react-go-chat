package websocket

import (
	"log"
	"net/http"

	"github.com/davidPardoC/go-chat/internal/chat/model"
	"github.com/davidPardoC/go-chat/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func checkOrigin(r *http.Request) bool {
	values := r.URL.Query()
	token := values.Get("token")

	isValid, _ := utils.IsTokenValid(token)

	return isValid
}

var upgrader = websocket.Upgrader{
	CheckOrigin: checkOrigin,
}

func wsHandler(ctx *gin.Context, hub *model.Hub) {
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Print("upgrade error:", err)
		return
	}
	defer c.Close()
	client := &model.Client{Hub: hub, Conn: c, Send: make(chan []byte, 256)}
	client.Hub.Register <- client

	go client.Read()
	go client.Write()
}

func StartWebSocketServer(r *gin.Engine) {
	hub := model.NewHub()
	go hub.Run()
	r.GET("/ws", func(c *gin.Context) {
		wsHandler(c, hub)
	})
}
