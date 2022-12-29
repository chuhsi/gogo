package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

type connection struct {
	ws   *websocket.Conn
	sc   chan []byte
	data *Data
}

var user_list = []string{}

var wu = &websocket.Upgrader{
	ReadBufferSize:  512,
	WriteBufferSize: 512,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func myws(w http.ResponseWriter, r *http.Request) {
	ws, err := wu.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	c := &connection{
		sc:   make(chan []byte, 256),
		ws:   ws,
		data: &Data{},
	}
	h.r <- c
	go c.writer()
	c.reader()
	defer func() {

	}()
}
func (c *connection) writer() {
	for msg := range c.sc {
		c.ws.WriteMessage(websocket.TextMessage, msg)
	}
	c.ws.Close()
}

func (c *connection) reader() {
	for {
		_, msg, err := c.ws.ReadMessage()
		if err != nil {
			h.r <- c
			break
		}
		json.Unmarshal(msg, &c.data)
		switch c.data.Type {
		case "login":
		case "user":
		case "logout":
			c.data.Type = "logout"
			user_list = del(user_list, c.data.User)
			data_b, _ := json.Marshal(c.data)
			h.b <- data_b
			h.r <- c
		default:
			fmt.Println("===============default===============")
		}
	}
}
func del(slice []string, user string) []string {
	count := len(slice)
	if count == 0 {
		return slice
	}
	if count == 1 && slice[0] == user {
		return []string{}
	}
	var n_slice = []string{}
	for i := range slice {
		if slice[i] == user && i == count {
			return slice[:count]
		} else if slice[i] == user {
			n_slice = append(slice[:i], slice[i+1:]...)
			break
		}
	}
	fmt.Println(n_slice)
	return n_slice
}
