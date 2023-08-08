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
	"strconv"
	"testing"

	"github.com/gorilla/mux"
)

func TestCRUD(t *testing.T) {
	itemRepo := repositories.NewMemoryItemRepository()
	itemUsecase := usecases.NewItemUsecase(itemRepo)

	router := mux.NewRouter()
	router = routers.NewItemRouter(router, itemUsecase)

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

	var items []entities.Item
	err = json.NewDecoder(resp.Body).Decode(&items)

	if err != nil {
		t.Fatalf("Error decoding response: %v", err)
	}

	resp.Body.Close()

	// Update Item
	if len(items) > 0 {
		itemToUpdate := items[0]
		itemToUpdate.Name = "Updated Item"
		updateJSON, _ := json.Marshal(itemToUpdate)
		req, _ := http.NewRequest(http.MethodPut, server.URL+"/items/"+strconv.Itoa(itemToUpdate.ID), bytes.NewBuffer(updateJSON))
		req.Header.Set("Content-Type", "application/json")
		resp, err = http.DefaultClient.Do(req)

		if err != nil {
			t.Fatalf("Error updating item: %v", err)
		}
		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
		}

		resp.Body.Close()
	}

	// Delete Item
	if len(items) > 0 {
		itemToDelete := items[0]
		req, _ := http.NewRequest(http.MethodDelete, server.URL+"/items/"+strconv.Itoa(itemToDelete.ID), nil)

		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			t.Fatalf("Error deleting item: %v", err)
		}

		if resp.StatusCode != http.StatusOK {
			t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
		}

		resp.Body.Close()
	}
}
