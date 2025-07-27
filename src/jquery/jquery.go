package main


import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

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
		Numbers1 int `json:"number1"`
		Numbers2 int `json:"number2"`
	}

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {

		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	// Calculate the sum of the numbers
	sum = requestData.Numbers1 + requestData.Numbers2

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

func indexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("Serving index.html")
	log.Printf("Request URL: %s", r.URL.Path)

	var staticDir string = "static"
	var fileName string = "index.html"

	if r.URL.Path != "/" {
		log.Printf("Request is for a non-root path: %s", r.URL.Path)
		fileName = r.URL.Path[1:] // Remove leading slash
		if _, err := os.Stat(fmt.Sprintf("%s/%s", staticDir, fileName)); os.IsNotExist(err) {
			log.Printf("File %s does not exist in static directory", fileName)
			http.NotFound(w, r)
			return
		}
	} else {
		log.Printf("Request is for the root path, serving %s", fileName)
	}

	http.ServeFile(w, r, fmt.Sprintf("%s/%s", staticDir, fileName))
}

func main() {
	// Register the handlers
	http.HandleFunc("/api/v1/add", addHandler)
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