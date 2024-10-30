package websocket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func checkOrigin(r *http.Request) bool {
	values := r.URL.Query()
	token := values.Get("token")

	return token != ""
}

var upgrader = websocket.Upgrader{
	CheckOrigin: checkOrigin,
}

func wsHandler(ctx *gin.Context) {
	c, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, []byte("message mi bro"))
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func StartWebSocketServer(r *gin.Engine) {
	fmt.Println("Starting - WEBSOCKET")
	r.GET("/ws", func(c *gin.Context) {
		wsHandler(c)
	})
}
