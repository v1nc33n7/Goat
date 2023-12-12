package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func createRoom(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade connection error: %v", err)
		return
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		conn.WriteMessage(websocket.TextMessage, []byte("Please provide username."))
		conn.Close()
		return
	}

	hub := NewHub()
	go hub.Run()

	client := &Client{username: username, hub: hub, conn: conn}
	hub.register <- client
	go client.Run()

	_, ok := hubList[hub.id]
	if !ok {
		hubList[hub.id] = hub
		log.Printf("New Hub created: [%s]", hub.id)
	}
}

func addClient(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Upgrade connection error: %v", err)
		return
	}

	username := r.URL.Query().Get("username")
	if username == "" {
		conn.WriteMessage(websocket.TextMessage, []byte("Please provide username."))
		conn.Close()
		return
	}

	hubId := r.URL.Query().Get("roomid")
	hub, ok := hubList[hubId]
	if !ok {
		conn.WriteMessage(websocket.TextMessage, []byte("Couldn't find room with this id."))
		conn.Close()
		return
	}

	client := &Client{username: username, hub: hub, conn: conn}
	hub.register <- client
	go client.Run()
}

func main() {
	hubList = make(map[string]*Hub)

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/create", createRoom)
	http.HandleFunc("/add", addClient)

	log.Printf("Listen on addr: %s\n", ":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
