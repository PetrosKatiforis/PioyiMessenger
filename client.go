package main

import (
    "github.com/gorilla/websocket"
)

// Builds upon the websocket connection with the addition of some state
// A nice way to abstract the connection and save the client's username
type Client struct {
    connection  *websocket.Conn
    username    string
}

// Creates a new client struct instance
func NewClient(connection *websocket.Conn, username string) *Client {
    return &Client {
        connection: connection,
        username: username,
    }
}

// Creates a new message and assigns the client as the author
func (c *Client) ComposeMessage(content string) Message {
    return Message {
        Author: c.username,
        Content: content,
    }
}

// Sending marshaled json to the client
func (c *Client) ReceiveJSON(data interface{}) {
    c.connection.WriteJSON(data)
}
