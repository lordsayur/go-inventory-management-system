package routers

import (
	"ims/api/controllers"
	"ims/core/interfaces"
	"net/http"

	"github.com/gorilla/mux"
)

func NewItemRouter(r *mux.Router, itemUsecase interfaces.ItemUsecase) *mux.Router {
	itemController := controllers.NewItemController(itemUsecase)

	r.HandleFunc("/items", enableCORS(itemController.CreateItem)).Methods("POST")
	r.HandleFunc("/items", enableCORS(itemController.GetAllItems)).Methods("GET")
	r.HandleFunc("/items/{id:[0-9]+}", enableCORS(itemController.GetById)).Methods("GET")
	r.HandleFunc("/items/{id:[0-9]+}", enableCORS(itemController.UpdateItem)).Methods("PUT")
	r.HandleFunc("/items/{id:[0-9]+}", enableCORS(itemController.DeleteItem)).Methods("DELETE")

	return r
}

func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Adjust this to your specific needs
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Handle preflight requests
		if r.Method == http.MethodOptions {
			return
		}

		next.ServeHTTP(w, r)
	}
}
