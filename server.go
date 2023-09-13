package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		// Allow any origin for demonstration purposes. In production, you should perform proper origin checks.
		return true
	},
}

func handlePing(ws *websocket.Conn) {
	defer ws.Close()

	for {
		// Read a message from the WebSocket connection
		_, msg, err := ws.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			return
		}

		fmt.Println(string(msg))

		// Respond with an "ok" message for any received message
		err = ws.WriteMessage(websocket.TextMessage, []byte("pong"))
		if err != nil {
			fmt.Println("Error writing 'pong' message:", err)
			return
		}
	}
}

func sendOkMessagesToClient(ws *websocket.Conn) {
	for {
		// time.Sleep(time.Second) // Send "ok" messages every seconds

		err := ws.WriteMessage(websocket.TextMessage, []byte("ok"))
		if err != nil {
			fmt.Println("Error writing 'ok' message:", err)
			return
		}
	}
}

func sendOkMessagesToClients(conns []*websocket.Conn) {
	for {
		time.Sleep(time.Second) // Send "ok" messages every second

		for _, ws := range conns {
			err := ws.WriteMessage(websocket.TextMessage, []byte("ok"))
			if err != nil {
				fmt.Println("Error writing 'ok' message:")
			}
		}
	}
}

func main() {
	var connections []*websocket.Conn

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		// Upgrade the HTTP connection to a WebSocket connection
		ws, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			fmt.Println("Error upgrading to WebSocket:", err)
			return
		}

		fmt.Println("Client Connected")

		connections = append(connections, ws)

		// Handle WebSocket "ping" messages
		go handlePing(ws)

		// Start a goroutine to send "ok" messages to all connected clients
		// go sendOkMessagesToClients(connections)

		// Start a goroutine to send "ok" messages to a new client
		// go sendOkMessagesToClient(ws)
	})

	// render html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	fmt.Println("WebSocket server listening on :8080")
	http.ListenAndServe(":8080", nil)
}
