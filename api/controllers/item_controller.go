package controllers

import (
	"encoding/json"
	"ims/core/entities"
	"ims/core/interfaces"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type ItemsController struct {
	itemUsecase interfaces.ItemUsecase
}

func NewItemController(itemUsecase interfaces.ItemUsecase) *ItemsController {
	return &ItemsController{
		itemUsecase: itemUsecase,
	}
}

func (h *ItemsController) CreateItem(w http.ResponseWriter, r *http.Request) {
	var newItem entities.Item
	err := json.NewDecoder(r.Body).Decode(&newItem)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.itemUsecase.CreateItem(newItem.Name)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ItemsController) GetAllItems(w http.ResponseWriter, r *http.Request) {
	sortField := r.URL.Query().Get("sortField")
	sortOrder := r.URL.Query().Get("sortOrder")

	items, err := h.itemUsecase.GetAllItems(sortField, sortOrder)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

func (h *ItemsController) UpdateItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	var updatedItem entities.Item
	err = json.NewDecoder(r.Body).Decode(&updatedItem)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedItem.ID = id
	err = h.itemUsecase.UpdateItem(&updatedItem)

	if err != nil {
		if err == entities.ErrNotFound {
			http.Error(w, "Item not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		return
	}
}

func (h *ItemsController) DeleteItem(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = h.itemUsecase.DeleteItem(id)

	if err != nil {
		if err == entities.ErrNotFound {
			http.Error(w, "Item not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}
}
