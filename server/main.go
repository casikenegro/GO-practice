package main

import (
	"fmt"
	"net/http"
)

func main() {
	server := NewServer(":3000")
	server.Handle("/", http.MethodGet, HandleRoot)
	server.Handle("/user", http.MethodPost, CreateUser)
	server.Handle("/api", http.MethodGet, server.addMiddleware(HandleHome, CheckAuth(), TestMiddleware()))
	fmt.Println("server running")
	server.listen()
}
