package main

import (
    "log"
    "fmt"
    "net/http"

    "github.com/gorilla/websocket"
)

// Allows the update of a regular http connection to a websocket
var upgrader = websocket.Upgrader{}

// Handles websocket.Conn objects and communicates with clients
type Server struct {
    // Storing all server clients to allow broadcasting
    clients map[*Client]bool

    // Useful for when a new client connects
    // And wants access to the current message history
    messageHistory []Message

    // Maximum limit of message history storage
    // If it passes the limit, it will pop out the oldest message
    historyLimit int
}

// Create a new server
func NewServer() *Server {
    return &Server {
        clients:        make(map[*Client]bool),
        messageHistory: make([]Message, 0),
        
        historyLimit:   100,
    }
}

// Handles incoming websocket connections
func (s *Server) HandleConnection(w http.ResponseWriter, r *http.Request) {
    // Client username will be found in the url parameters of the request
    usernames, ok := r.URL.Query()["username"]

    // Check if the username is valid
    if !ok || usernames[0] == "Server" {
        log.Println("Username is invalid or missing!")
        return 
    }

    // Try to upgrade the connection to a websocket
    connection, err := upgrader.Upgrade(w, r, nil)

    if err != nil {
        log.Println("Failed to upgrade http connection!", err)
        return
    }

    client := NewClient(connection, usernames[0])

    // Storing the connection into the map
    s.clients[client] = true

    defer connection.Close()
    defer delete(s.clients, client)

    s.HandleLoad(client)

    for {
        _, message, err := connection.ReadMessage()

        // If an error occurs, then the client has been disconnected
        if err != nil {
            s.HandleDisconnect(client)
            break
        }

        go s.HandleMessage(client, string(message))
    }
}

// Executes when a new connection is established
// Provides the message history to the new client and broadcasts a new welcome message
func (s *Server) HandleLoad(client *Client) {
    // Send message history to client
    client.ReceiveJSON(s.messageHistory)

    message := Message {
        Author: "Server",
        Content: fmt.Sprintf("A new user has entered the chat. Please welcome %s!", client.username),
    }

    s.UpdateHistory(message)

    s.BroadcastJSON(&message)
}

// Process a client-sent message and broadcasts it to the chat
func (s *Server) HandleMessage(client *Client, message string) {
    // If the message is empty, don't send anything
    if len(message) == 0 {
        return
    }

    data := client.ComposeMessage(message)

    // Add message to the chat history
    s.UpdateHistory(data)

    s.BroadcastJSON(&data)
}

// Handles the disconnection of a client
// Will send a server message indicating that the user has left the chat
func (s *Server) HandleDisconnect(client *Client) {
    message := Message {
        Author: "Server",
        Content: fmt.Sprintf("%s has left the chat!", client.username),
    }

    s.UpdateHistory(message)

    s.BroadcastJSON(&message)
}

// Broadcasts JSON data to all server clients
func (s *Server) BroadcastJSON(data interface{}) {
    for client, _ := range(s.clients) {
        
        client.ReceiveJSON(data)
    }
}

// Push message to server history and manage storage limit
func (s *Server) UpdateHistory(message Message) {
    
    if len(s.messageHistory) > s.historyLimit {
        
        // Remove oldest message
        s.messageHistory = s.messageHistory[1:]
    }

    s.messageHistory = append(s.messageHistory, message)
}


