package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Client struct {
	username string
	hub      *Hub
	conn     *websocket.Conn
}

func validMsg(msg []byte) bool {
	return string(msg) == ""
}

func (c *Client) Run() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		if validMsg(msg) {
			continue
		}
		c.hub.broadcast <- fmt.Sprintf("%s: %s", c.username, msg)
	}
}
