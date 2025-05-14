package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"os"
	"os/signal"
	"strings"
	"time"
)

var serverAddr = "ws://localhost:8080/ws"

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Get user's name
	fmt.Print("Enter your name: ")
	nameRaw, _ := reader.ReadString('\n')
	name := strings.TrimSpace(nameRaw)

	// Connect to WebSocket server
	conn, _, err := websocket.DefaultDialer.Dial(serverAddr, nil)
	if err != nil {
		log.Fatal("Connection error:", err)
	}
	defer conn.Close()

	// Send name as the first message
	err = conn.WriteMessage(websocket.TextMessage, []byte(name))
	if err != nil {
		log.Fatal("Failed to send name:", err)
	}

	// Interrupt signal handling
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// Channel to receive all messages
	messageChan := make(chan string, 100)

	// Goroutine to receive messages
	go func() {
		for {
			_, msg, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("\n[Disconnected from server]")
				os.Exit(0)
			}
			messageChan <- string(msg)
		}
	}()

	// Menu loop
	for {
		fmt.Println("\n--- Menu ---")
		fmt.Println("1. Exit")
		fmt.Println("2. Show Messages (press 'q' + Enter to stop)")
		fmt.Println("3. Send Message")
		fmt.Print("Choose option: ")

		optionRaw, _ := reader.ReadString('\n')
		option := strings.TrimSpace(optionRaw)

		switch option {
		case "1":
			fmt.Println("Goodbye!")
			return

		case "2":
			fmt.Println("Listening to messages (press 'q' + Enter to stop)...")
			done := make(chan bool)

			// Goroutine to print messages while in this mode
			go func() {
				for {
					select {
					case msg := <-messageChan:
						// Show all messages except user's own
						if !strings.HasPrefix(msg, name+":") {
							fmt.Println(msg)
						}
					case <-done:
						return
					}
				}
			}()

			// Wait for 'q' to quit showing messages
			for {
				text, _ := reader.ReadString('\n')
				if strings.TrimSpace(text) == "q" {
					done <- true
					break
				}
			}

		case "3":
			fmt.Print("Enter message: ")
			msgRaw, _ := reader.ReadString('\n')
			msg := strings.TrimSpace(msgRaw)
			if msg != "" {
				err := conn.WriteMessage(websocket.TextMessage, []byte(msg))
				if err != nil {
					log.Println("Write error:", err)
				}
			} else {
				fmt.Println("Empty message. Not sent.")
			}

		default:
			fmt.Println("Invalid option.")
		}

		time.Sleep(200 * time.Millisecond)
	}
}
