package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrager = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	http.HandleFunc("/ec", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrager.Upgrade(w, r, nil)
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				return
			}
			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(p))

			if err = conn.WriteMessage(messageType, p); err != nil {
				return
			}
		}
	})

	http.ListenAndServe(":9090", nil)
}
