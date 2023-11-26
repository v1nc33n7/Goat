package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

type Message struct {
	username string
	msgType  int
	msg      []byte
}

type Client struct {
	username string
	hub      *Hub
	conn     *websocket.Conn
}

func NewMessage(msg string) Message {
	return Message{
		msgType: websocket.TextMessage,
		msg:     []byte(msg),
	}
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
		c.hub.broadcast <- NewMessage(fmt.Sprintf("%s: %s", c.username, msg))
	}
}
