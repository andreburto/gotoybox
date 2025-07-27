package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/file/", FileHandler) // TODO: Why does this need a trailing slash?
	http.HandleFunc("/api/v1/list", ListHandlers)
	http.HandleFunc("/", IndexHandler)

	// TODO: Get port from environment variable or config
	port := ":8080"
	fmt.Printf("Server is running on port %s\n", port)
	if err := http.ListenAndServe(port, nil); err != nil {
		log.Fatalf("Could not start server: %s\n", err)
	}
}
