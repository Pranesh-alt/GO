package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

type Client struct {
	conn *websocket.Conn
	name string
}

type Server struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan string
	mu         sync.Mutex
}

func NewServer() *Server {
	return &Server{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan string),
	}
}

func (s *Server) Start() {
	for {
		select {
		case client := <-s.register:
			s.mu.Lock()
			s.clients[client] = true
			s.mu.Unlock()
			log.Printf("Client %s joined the chat", client.name)

		case client := <-s.unregister:
			s.mu.Lock()
			delete(s.clients, client)
			s.mu.Unlock()
			log.Printf("Client %s left the chat", client.name)

		case message := <-s.broadcast:
			// Print the message to the server's console to see the chats
			fmt.Println("Server Chat Message: ", message)

			// Broadcast the message to all connected clients
			s.mu.Lock()
			for client := range s.clients {
				err := client.conn.WriteMessage(websocket.TextMessage, []byte(message))
				if err != nil {
					log.Println("Error sending message:", err)
					client.conn.Close()
					delete(s.clients, client)
				}
			}
			s.mu.Unlock()
		}
	}
}

func (s *Server) handleConnection(conn *websocket.Conn) {
	defer conn.Close()

	// Get the client's name
	conn.WriteMessage(websocket.TextMessage, []byte("Enter your name: "))
	_, nameMsg, _ := conn.ReadMessage()
	clientName := string(nameMsg)

	client := &Client{conn: conn, name: clientName}
	s.register <- client

	// Listen for messages from this client
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				s.unregister <- client
				log.Printf("Client %s disconnected", clientName)
				break
			}
			// Broadcast the message to all clients
			message := fmt.Sprintf("%s: %s", clientName, msg)
			s.broadcast <- message
		}
	}()

	// Keep the client connected
	for {
		time.Sleep(time.Second)
	}
}

func main() {
	server := NewServer()
	go server.Start()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		upgrader := websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool { return true },
		}
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Println("Failed to upgrade WebSocket:", err)
			return
		}
		server.handleConnection(conn)
	})

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
