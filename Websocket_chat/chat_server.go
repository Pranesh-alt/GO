package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

type Client struct {
	conn *websocket.Conn
	send chan []byte
	name string
}

type Server struct {
	clients    map[*Client]bool
	register   chan *Client
	unregister chan *Client
	broadcast  chan []byte
	mu         sync.Mutex
	db         *sql.DB
}

func NewServer(db *sql.DB) *Server {
	return &Server{
		clients:    make(map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan []byte),
		db:         db,
	}
}

func (s *Server) Run() {
	for {
		select {
		case client := <-s.register:
			s.mu.Lock()
			s.clients[client] = true
			s.mu.Unlock()
			fmt.Printf("[JOIN] %s joined\n", client.name)

		case client := <-s.unregister:
			s.mu.Lock()
			if _, ok := s.clients[client]; ok {
				delete(s.clients, client)
				close(client.send)
				fmt.Printf("[LEAVE] %s left\n", client.name)
			}
			s.mu.Unlock()

		case message := <-s.broadcast:
			s.mu.Lock()
			for client := range s.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(s.clients, client)
				}
			}
			s.mu.Unlock()
		}
	}
}

func (c *Client) readPump(s *Server) {
	defer func() {
		s.unregister <- c
		c.conn.Close()
	}()

	for {
		_, msg, err := c.conn.ReadMessage()
		if err != nil {
			break
		}
		text := fmt.Sprintf("%s: %s", c.name, msg)
		s.broadcast <- []byte(text)

		// Save message to DB
		_, err = s.db.Exec("INSERT INTO messages (username, content) VALUES (?, ?)", c.name, string(msg))
		if err != nil {
			fmt.Println("DB insert error:", err)
		}
	}
}

func (c *Client) writePump() {
	for msg := range c.send {
		err := c.conn.WriteMessage(websocket.TextMessage, msg)
		if err != nil {
			break
		}
	}
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func serveWs(s *Server, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}

	// Get name as the first message
	_, nameMsg, err := conn.ReadMessage()
	if err != nil {
		conn.Close()
		return
	}
	name := string(nameMsg)

	client := &Client{
		conn: conn,
		send: make(chan []byte, 256),
		name: name,
	}

	// Send last 10 messages from DB
	rows, err := s.db.Query("SELECT username, content FROM messages ORDER BY timestamp DESC LIMIT 10")
	if err == nil {
		defer rows.Close()
		var username, content string
		var history []string
		for rows.Next() {
			rows.Scan(&username, &content)
			history = append([]string{fmt.Sprintf("%s: %s", username, content)}, history...)
		}
		for _, msg := range history {
			client.send <- []byte(msg)
		}
	}

	s.register <- client

	go client.writePump()
	go client.readPump(s)
}

func initDB() (*sql.DB, error) {
	// Update with your own credentials
	dsn := "root:62145090@tcp(127.0.0.1:3306)/chatapp"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return db, nil
}

func main() {
	db, err := initDB()
	if err != nil {
		fmt.Println("Database connection error:", err)
		return
	}
	defer db.Close()

	server := NewServer(db)
	go server.Run()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(server, w, r)
	})

	fmt.Println("Server started on ws://localhost:8080/ws")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Server error:", err)
	}
}
