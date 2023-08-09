package main

import (
	"fmt"
	"ims/api/routers"
	"ims/core/usecases"
	"infrastructure/repositories"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// itemRepository := repositories.NewMemoryItemRepository()
	itemRepository := repositories.NewXORMItemRepository()

	itemUsecase := usecases.NewItemUsecase(itemRepository)

	router := mux.NewRouter()
	router = routers.NewItemRouter(router, itemUsecase)

	port := 8080
	fmt.Printf("Server started on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
