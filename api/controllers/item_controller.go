package controllers

import (
	"encoding/json"
	"ims/core/entities"
	"ims/core/usecases"
	"net/http"
)

type ItemsController struct {
	itemUsecase usecases.ItemUsecase
}

func NewItemController(itemUsecase usecases.ItemUsecase) *ItemsController {
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
	items, err := h.itemUsecase.GetAllItems()

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
