package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var CheckFile string = "check.txt"

func CheckFileExists() bool {
	// Check if the file exists
	if _, err := os.Stat(CheckFile); os.IsNotExist(err) {
		// log.Printf("File %s does not exist", CheckFile)
		return false
	}
	log.Printf("File %s exists", CheckFile)
	return true
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

// handle the /api/v1/waiter endpoint
func waiterHandler(w http.ResponseWriter, r *http.Request) {
	var fileExists string = "no"

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	// Check if the request method is one of the allowed methods
	allowedMethods := []string{http.MethodGet, http.MethodPost, http.MethodDelete}
	isAllowed := false
	for _, method := range allowedMethods {
		if r.Method == method {
			isAllowed = true
			break
		}
	}

	if !isAllowed {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodPost {
		// Attempt to create the file
		file, err := os.Create(CheckFile)
		if err != nil {
			var error_msg string = fmt.Sprintf("Error creating file %s: %v", CheckFile, err)
			// log.Println(error_msg)
			http.Error(w, error_msg, http.StatusInternalServerError)
			return
		}
		defer file.Close()
	}
	
	if r.Method == http.MethodDelete {
		// Attempt to delete the file
		err := os.Remove(CheckFile)
		if err != nil {
			var error_msg string = fmt.Sprintf("Error deleting file %s: %v", CheckFile, err)
			// log.Println(error_msg)
			http.Error(w, error_msg, http.StatusInternalServerError)
			return
		}
	}
	
	// Check if the file exists
	if CheckFileExists() {
		fileExists = "yes"
	}

	// Respond with a simple message
	response := fmt.Sprintf(`{"fileExists": "%s"}`, fileExists)
	
	_, err := w.Write([]byte(response))
	if err != nil {
		// log.Printf("Error writing response: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Register the handler for the /api/v1/waiter endpoint
	http.HandleFunc("/api/v1/waiter", waiterHandler)
	http.HandleFunc("/", indexHandler)

	// Get the port from the environment variable or use a default
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000" // Default port
	}

	srv := &http.Server{
		Addr: fmt.Sprintf(":%s", port),
		Handler: nil,
	}

	// Start the server
	log.Fatal(srv.ListenAndServe())
}