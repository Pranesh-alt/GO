package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

var (
	clients     = make(map[*websocket.Conn]bool) // Connected clients
	clientsLock sync.Mutex                       // Mutex to handle concurrent access to the clients map
)

// WebSocket handler function
func handleConnections(w http.ResponseWriter, r *http.Request) {
	// Upgrade HTTP connection to WebSocket
	conn, err := websocket.Upgrade(w, r, nil, 1024, 1024)
	if err != nil {
		fmt.Println("Error upgrading connection:", err)
		return
	}
	defer conn.Close()

	// Add the new client to the map (with mutex to ensure thread safety)
	clientsLock.Lock()
	clients[conn] = true
	clientsLock.Unlock()

	// Listen for incoming messages from this client
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		// Broadcast the message to all connected clients (with mutex to ensure thread safety)
		clientsLock.Lock()
		for client := range clients {
			if err := client.WriteMessage(messageType, message); err != nil {
				fmt.Println("Error sending message:", err)
			}
		}
		clientsLock.Unlock()
	}

	// Remove the client from the map once disconnected (with mutex to ensure thread safety)
	clientsLock.Lock()
	delete(clients, conn)
	clientsLock.Unlock()
}

func main() {
	// Serve WebSocket connections on "/ws"
	http.HandleFunc("/ws", handleConnections)

	// Start the server
	fmt.Println("Server started on ws://localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
