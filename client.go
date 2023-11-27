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
		c.hub.broadcast <- fmt.Sprintf("%s: %s", c.username, msg)
	}
}
