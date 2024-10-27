package webscoket

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func checkOrigin(r *http.Request) bool {
	fmt.Println("Authorizing")
	if r.Header.Get("authorization") == "" {
		return false
	}
	return true
}

var upgrader = websocket.Upgrader{
	CheckOrigin: checkOrigin,
}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
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

func StartWebSocketServer() {
	http.HandleFunc("/ws", wsHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
