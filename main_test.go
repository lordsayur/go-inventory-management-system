package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCRUD(t *testing.T) {
	r := mux.NewRouter()
	r.HandleFunc("/items", createItem).Methods("POST")
	r.HandleFunc("/items", readItems).Methods("GET")

	server := httptest.NewServer(r)
	defer server.Close()

	// Create Item
	newItem := Item{Name: "Test Item"}
	newItemJSON, _ := json.Marshal(newItem)

	resp, err := http.Post(server.URL+"/items", "application/json", bytes.NewBuffer(newItemJSON))

	if err != nil {
		t.Fatalf("Error creating item: %v", err)
	}

	if resp.StatusCode != http.StatusCreated {
		t.Errorf("Expected status code %d but got %d", http.StatusCreated, resp.StatusCode)
	}

	resp.Body.Close()

	// Read Items
	resp, err = http.Get(server.URL + "/items")

	if err != nil {
		t.Fatalf("Error reading items: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	resp.Body.Close()
}
