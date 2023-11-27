package main

import (
	"fmt"
	"log"
	"math/rand"

	"github.com/gorilla/websocket"
)

type HubManager struct {
	register   chan *Hub
	unregister chan *Hub
	hubs       map[string]*Hub
}

type Hub struct {
	id         string
	hm         *HubManager
	register   chan *Client
	unregister chan *Client
	broadcast  chan string
	clients    map[*Client]bool
}

var lettersRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randomId(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = lettersRunes[rand.Intn(len(lettersRunes))]
	}
	return string(b)
}

func NewManager() *HubManager {
	return &HubManager{
		register:   make(chan *Hub),
		unregister: make(chan *Hub),
		hubs:       make(map[string]*Hub),
	}
}

func NewHub(hm *HubManager) *Hub {
	return &Hub{
		id:         randomId(8),
		hm:         hm,
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan string),
		clients:    make(map[*Client]bool),
	}
}

func (hm *HubManager) Run() {
	for {
		select {
		case hub := <-hm.register:
			_, ok := hm.hubs[hub.id]
			if !ok {
				hm.hubs[hub.id] = hub

				log.Printf("New Hub created: [%s]", hub.id)
			}
		}
	}
}

func (h *Hub) runMessage() {
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
	go h.runMessage()

	for {
		select {
		case client := <-h.register:
			_, ok := h.clients[client]
			if !ok {
				h.clients[client] = true
				h.broadcast <- fmt.Sprintf("Room [%s]: Welcome %s", h.id, client.username)

				log.Printf("Register Client[%s] -> Hub[%s]", client.username, h.id)
			}
		case client := <-h.unregister:
			_, ok := h.clients[client]
			if ok {
				delete(h.clients, client)
				client.conn.Close()
				h.broadcast <- fmt.Sprintf("Room [%s]: Bye %s", h.id, client.username)

				log.Printf("Unregister Client[%s] <- Hub[%s]", client.username, h.id)
			}
			if len(h.clients) == 0 {
				delete(h.hm.hubs, h.id)

				log.Printf("Removed Hub[%s]", h.id)
			}
		}
	}
}
