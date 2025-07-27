package main


import (
	// "encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
)

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
	// Register the handler for the /api/v1/add endpoint
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