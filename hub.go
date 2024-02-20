package main

import (
	"fmt"
	"log"

	"github.com/gorilla/websocket"
)

var hubList map[string]*Hub

type Hub struct {
	id         string
	register   chan *Client
	unregister chan *Client
	broadcast  chan string
	clients    map[*Client]bool
}

func NewHub(id string) *Hub {
	hub := &Hub{
		id:         id,
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan string),
		clients:    make(map[*Client]bool),
	}

	hubList[hub.id] = hub
	log.Printf("New Hub created: [%s]", hub.id)

	return hub
}

func (h *Hub) broadcastRun() {
	for {
		message := <-h.broadcast
		for client := range h.clients {
			err := client.conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				h.unregister <- client
				client.conn.Close()

				log.Println("Error, write to clients:", err)
			}
		}
	}
}

func (h *Hub) Run() {
	go h.broadcastRun()

	for {
		select {
		case client := <-h.register:
			_, ok := h.clients[client]
			if !ok {
				h.clients[client] = true
				h.broadcast <- fmt.Sprintf("Room [%s]: Hello %s!", h.id, client.username)

				log.Printf("Register Client[%s] -> Hub[%s]", client.username, h.id)
			}
		case client := <-h.unregister:
			_, ok := h.clients[client]
			if ok {
				delete(h.clients, client)
				client.conn.Close()

				h.broadcast <- fmt.Sprintf("Room [%s]: Goodbye %s!", h.id, client.username)

				log.Printf("Unregister Client[%s] <- Hub[%s]", client.username, h.id)
			}
			if len(h.clients) == 0 {
				delete(hubList, h.id)

				log.Printf("Removed Hub[%s]", h.id)
			}
		}
	}
}
