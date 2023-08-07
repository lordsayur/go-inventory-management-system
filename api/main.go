package main

import (
	"encoding/json"
	"fmt"
	"ims/entity"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var items = make(map[int]entity.Item)
var lastID = 0

func createItem(w http.ResponseWriter, r *http.Request) {
	var newItem entity.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	lastID++
	newItem.ID = lastID
	items[newItem.ID] = newItem

	w.WriteHeader(http.StatusCreated)
}

func readItems(w http.ResponseWriter, r *http.Request) {
	itemList := []entity.Item{}

	for _, item := range items {
		itemList = append(itemList, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(itemList)
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/items", createItem).Methods("POST")
	r.HandleFunc("/items", readItems).Methods("GET")

	port := 8080
	fmt.Printf("Server started on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), r))
}
