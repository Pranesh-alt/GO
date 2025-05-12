package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
)

var wsURL = "ws://localhost:8080/ws"

func main() {
	// Connect to the WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		fmt.Println("Error connecting to WebSocket server:", err)
		return
	}
	defer conn.Close()

	// Read the user's name
	var userName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&userName)

	// Send username to server
	conn.WriteMessage(websocket.TextMessage, []byte(userName+" has joined the chat"))

	// Goroutine to receive messages from the server
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				return
			}
			// Print received message to console
			fmt.Println(string(message))
		}
	}()

	// Main loop for sending messages
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter your message: ")
		if !scanner.Scan() {
			fmt.Println("Input error.")
			break
		}
		message := scanner.Text()

		// Send message to WebSocket server
		conn.WriteMessage(websocket.TextMessage, []byte(userName+": "+message))
	}
}
