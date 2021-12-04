package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"

	"prw_server/app/pkg/handlers"
)

func main() {
	handler := handlers.NewHandler()

	router := chi.NewRouter()

	router.Get("/hello", handler.Hello)

	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}
