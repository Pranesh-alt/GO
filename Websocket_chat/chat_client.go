package main

import (
	"bufio"
	"fmt"
	"github.com/gorilla/websocket"
	"os"
	"os/exec"
)

var wsURL = "ws://localhost:8080/ws"

func clearScreen() {
	cmd := exec.Command("clear")
	if err := cmd.Run(); err != nil {
		fmt.Println("Error clearing screen:", err)
	}
}

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	if err != nil {
		fmt.Println("Error connecting to WebSocket server:", err)
		return
	}
	defer conn.Close()

	// Read user's name
	var userName string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&userName)

	// Send the name to the server
	conn.WriteMessage(websocket.TextMessage, []byte(userName))

	// Handle user input for sending messages
	go func() {
		for {
			_, message, err := conn.ReadMessage()
			if err != nil {
				fmt.Println("Error reading message:", err)
				return
			}
			fmt.Println(message)
		}
	}()

	// Menu loop
	scanner := bufio.NewScanner(os.Stdin)
	for {
		clearScreen()
		fmt.Println("Chat Menu:")
		fmt.Println("1. Exit")
		fmt.Println("2. Show Messages (Press 'q' and hit enter to go back)")
		fmt.Println("3. Send Message")
		fmt.Print("Choose an option: ")

		scanner.Scan()
		option := scanner.Text()

		switch option {
		case "1":
			conn.WriteMessage(websocket.TextMessage, []byte(userName+" has left the chat"))
			return
		case "2":
			// Just wait for the message feed to come through the go routine
			fmt.Println("Press 'q' to return to the menu.")
			for {
				if scanner.Scan() {
					if scanner.Text() == "q" {
						break
					}
				}
			}
		case "3":
			fmt.Print("Enter your message: ")
			scanner.Scan()
			message := scanner.Text()
			conn.WriteMessage(websocket.TextMessage, []byte(message))
		}
	}
}
