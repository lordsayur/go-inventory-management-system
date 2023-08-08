package routers

import (
	"ims/api/controllers"
	"ims/core/interfaces"

	"github.com/gorilla/mux"
)

func NewItemRouter(r *mux.Router, itemUsecase interfaces.ItemUsecase) *mux.Router {
	itemController := controllers.NewItemController(itemUsecase)

	r.HandleFunc("/items", itemController.CreateItem).Methods("POST")
	r.HandleFunc("/items", itemController.GetAllItems).Methods("GET")

	return r
}
