package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

func FileHandler(w http.ResponseWriter, r *http.Request) {
	var filePath string = r.URL.Path[len("/api/v1/file/"):]
	
	w.Header().Set("Content-Type", "text/plain")

	log.Printf("FileHandler called with filePath: %s", filePath)
	log.Printf("Request Method: %s", r.Method)

	if r.Method == http.MethodPut {
		body, err := io.ReadAll(r.Body)
		if err != nil {
			log.Printf("Error reading request body: %v", err)
			http.Error(w, "Bad request", http.StatusBadRequest)
			return
		}
		var doc string = string(body)

		SaveFileToDisk(doc, filePath)

		_, err2 := w.Write([]byte("OK"))
		if err2 != nil {
			log.Printf("Error writing response: %v", err2)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		}
	} else if r.Method == http.MethodGet {
		w.Write([]byte(LoadFileFromDisk(filePath)))
	} else {
		http.Error(w, "Unsupported method", http.StatusMethodNotAllowed)
	}
}

func ListHandlers(w http.ResponseWriter, r *http.Request) {
	log.Println("ListHandlers called")
	log.Printf("Request URL: %s", r.URL.Path)

	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var paths []string

	fileMap := LoadFileMap()
	for path, _ := range fileMap {
		paths = append(paths, path)
	}

	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, "{\"paths\": [\"%s\"]}", strings.Join(paths, "\", \""))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("IndexHandler called")
	log.Printf("Request URL: %s", r.URL.Path)
	fmt.Fprintln(w, "Welcome to the Drive API")
}