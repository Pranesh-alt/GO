package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var serverAddr = "ws://localhost:8080/ws"

func main() {
	fmt.Print("Enter your name: ")
	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n')
	name = name[:len(name)-1]

	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer conn.Close()

	// Send name as first message
	conn.WriteMessage(websocket.TextMessage, []byte(name))

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	showMessages := false

	// Start message receiver
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Disconnected from server.")
				os.Exit(0)
			}
			if showMessages {
				fmt.Println(string(msg))
			}
		}
	}()

	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Exit")
		fmt.Println("2. show-messages (press 'q' + Enter to return)")
		fmt.Println("3. send-message")
		fmt.Print("Choose option: ")

		option, _ := reader.ReadString('\n')

		switch option {
		case "1\n":
			fmt.Println("Goodbye!")
			return

		case "2\n":
			showMessages = true
			fmt.Println("Listening to messages (press 'q' then Enter to quit)...")
			for {
				text, _ := reader.ReadString('\n')
				if text == "q\n" {
					showMessages = false
					break
				}
			}

		case "3\n":
			fmt.Print("Enter message: ")
			message, _ := reader.ReadString('\n')
			err := conn.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("Write error:", err)
			}

		default:
			fmt.Println("Invalid option.")
		}

		time.Sleep(200 * time.Millisecond)
	}
}
