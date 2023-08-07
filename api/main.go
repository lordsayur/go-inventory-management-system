package main

import (
	"fmt"
	"ims/api/routers"
	"ims/core/usecases"
	"infrastructure/repositories"
	"log"
	"net/http"
)

func main() {
	itemRepo := repositories.NewMemoryItemRepository()
	itemUsecase := usecases.NewItemUsecase(itemRepo)
	router := routers.NewRouter(*itemUsecase)

	port := 8080
	fmt.Printf("Server started on port %d...\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), router))
}
