package main 

// Stores all the necessary info for a chat message
// Will be sent as json throughout the websocket connection
type Message struct {
    Author  string `json:"author"`
    Content string `json:"content"`
}
