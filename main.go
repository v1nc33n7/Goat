package main

import (
	"log"
	"math/rand"
	"net/http"

	"github.com/gorilla/websocket"
)

const addr = "127.0.0.1:8080"

const lenghtId = 8

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var lettersRunes = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

func randomId(str string) string {
	_, ok := hubList[str]
	if !ok && str != "" {
		return str
	}

	b := make([]rune, lenghtId)
	for i := range b {
		b[i] = lettersRunes[rand.Intn(len(lettersRunes))]
	}
	return randomId(string(b))
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

	hub := NewHub(randomId(""))
	go hub.Run()

	client := &Client{username: username, hub: hub, conn: conn}
	hub.register <- client
	go client.Run()
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

	http.Handle("/", http.FileServer(http.Dir("./public")))
	http.HandleFunc("/create", createRoom)
	http.HandleFunc("/add", addClient)

	log.Printf("Listen on addr: %s\n", addr)
	log.Fatal(http.ListenAndServe(addr, nil))
}
