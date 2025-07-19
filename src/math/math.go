package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/index.html")
}

// handle the /api/v1/add endpoint
func addHandler(w http.ResponseWriter, r *http.Request) {
	var response string
	var sum int = 0

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON body
	var requestData struct {
		Numbers []int `json:"numbers"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {

		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Calculate the sum of the numbers
	for _, num := range requestData.Numbers {
		sum += num
	}

	// Create the response
	response = fmt.Sprintf(`{"result": %d}`, sum)

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	log.Printf("Add Response: %s", response)

	_, err2 := w.Write([]byte(response))
	if err2 != nil {
		log.Printf("Error writing response: %v", err2)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func subtractHandler(w http.ResponseWriter, r *http.Request) {
	var response string
	var result int = 0

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the JSON body
	var requestData struct {
		Number1 int `json:"numbers1"`
		Number2 int `json:"numbers2"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	// Initialize result with the first number
	result = requestData.Number1 - requestData.Number2

	// Create the response
	response = fmt.Sprintf(`{"result": %d}`, result)

	// Set the content type to JSON
	w.Header().Set("Content-Type", "application/json")

	log.Printf("Subtract Response: %s", response)

	_, err2 := w.Write([]byte(response))
	if err2 != nil {
		log.Printf("Error writing response: %v", err2)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	// Register the handler for the /api/v1/add endpoint
	http.HandleFunc("/api/v1/add", addHandler)
	http.HandleFunc("/api/v1/subtract", subtractHandler)
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