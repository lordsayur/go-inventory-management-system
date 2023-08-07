package main

import (
	"bytes"
	"encoding/json"
	"ims/api/routers"
	"ims/core/entities"
	"ims/core/usecases"
	"infrastructure/repositories"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestCRUD(t *testing.T) {
	itemRepo := repositories.NewMemoryItemRepository()
	itemUsecase := usecases.NewItemUsecase(itemRepo)
	router := routers.NewRouter(*itemUsecase)

	server := httptest.NewServer(router)
	defer server.Close()

	// Create Item
	newItem := entities.Item{Name: "Test Item"}
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
