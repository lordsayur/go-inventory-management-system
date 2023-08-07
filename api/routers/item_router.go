package routers

import (
	"ims/api/controllers"
	"ims/core/usecases"

	"github.com/gorilla/mux"
)

func NewRouter(itemUsecase usecases.ItemUsecase) *mux.Router {
	r := mux.NewRouter()

	itemController := controllers.NewItemController(itemUsecase)

	r.HandleFunc("/items", itemController.CreateItem).Methods("POST")
	r.HandleFunc("/items", itemController.GetAllItems).Methods("GET")

	return r
}
