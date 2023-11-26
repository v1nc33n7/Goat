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

func createRoom(hubManager *HubManager, w http.ResponseWriter, r *http.Request) {
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

	hub := NewHub(hubManager)
	go hub.Run()

	client := &Client{username: username, hub: hub, conn: conn}
	hub.register <- client
	go client.Run()

	hubManager.register <- hub
}

func addClient(hubManager *HubManager, w http.ResponseWriter, r *http.Request) {
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
	hub, ok := hubManager.hubs[hubId]
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
	hubManager := NewManager()
	go hubManager.Run()

	http.Handle("/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/create", func(w http.ResponseWriter, r *http.Request) {
		createRoom(hubManager, w, r)
	})
	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		addClient(hubManager, w, r)
	})

	log.Printf("Listen on addr: %s\n", ":8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
