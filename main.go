package main

import (
    "log"
    "net/http"
)

func main() {
    server := NewServer()

    // Endpoint in which the websocket connections will be created and handled
    http.HandleFunc("/server/socket", server.HandleConnection)
    
    // Serving the svelte frontend
    http.Handle("/", http.FileServer(http.Dir("public/build")))

    log.Println("Booting up the server!")

    // Boots up the server
    log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
