package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/ilyakaznacheev/cleanenv"

	"prw_server/app/config"
	"prw_server/app/pkg/handlers"
)

func main() {
	cfg := config.Server{}
	err := cleanenv.ReadConfig("config.yml", &cfg)
	if err != nil {
		log.Fatal(err)
	}
	handler := handlers.NewHandler()

	router := chi.NewRouter()

	router.Get("/hello", handler.Hello)

	path := cfg.Host + ":" + cfg.Port

	log.Print("Starting server")
	err = http.ListenAndServe(path, router)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Shutting server down")
}
