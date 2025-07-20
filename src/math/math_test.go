package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetupRequest(method, path string, body []byte, t *testing.T) *http.Request {
	req, err := http.NewRequest(method, path, bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("Failed to create request: %v", err)
	}
	req.Header.Set("Content-Type", "application/json")
	return req
}

func TestAdd1(t *testing.T) {
	// Create a sample JSON payload
	payload := map[string]interface{}{
		"numbers": []int{1, 2},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Create a new HTTP request with the JSON payload
	req := SetupRequest("POST", "/add", jsonData, t)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	addHandler(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Optionally, check the response body
	expected := `{"result": 3}` // Example expected response
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestAdd2(t *testing.T) {
	// Create a sample JSON payload
	payload := map[string]interface{}{
		"numbers": []int{5, 10, 15},
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Create a new HTTP request with the JSON payload
	req := SetupRequest("POST", "/add", jsonData, t)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	addHandler(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Optionally, check the response body
	expected := `{"result": 30}` // Example expected response
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestSubtract1(t *testing.T) {
	// Create a sample JSON payload
	payload := map[string]interface{}{
		"numbers1": 10,
		"numbers2": 5,
	}

	jsonData, err := json.Marshal(payload)
	if err != nil {
		t.Fatalf("Failed to marshal JSON: %v", err)
	}

	// Create a new HTTP request with the JSON payload
	req := SetupRequest("POST", "/subtract", jsonData, t)

	// Create a ResponseRecorder to capture the response
	rr := httptest.NewRecorder()

	// Call the handler function
	subtractHandler(rr, req)

	// Check the response status code
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Optionally, check the response body
	expected := `{"result": 5}` // Example expected response
	if rr.Body.String() != expected {
		t.Errorf("Handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}
