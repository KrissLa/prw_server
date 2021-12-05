package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"prw_server/app/pkg/api/jokes"
	"syscall"
	"time"

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

	apiClient := jokes.NewJokeClient(cfg.JokeURL)

	handler := handlers.NewHandler(apiClient)

	router := chi.NewRouter()

	router.Get("/hello", handler.Hello)

	path := cfg.Host + ":" + cfg.Port

	server := &http.Server{
		Addr: path,
		Handler: router,
	}

	quit := make(chan os.Signal, 1)
	done := make(chan error, 1)

	signal.Notify(quit, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-quit
		ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
		cancel()
		done <- server.Shutdown(ctx)
	}()

	log.Print("Starting server")
	_ = server.ListenAndServe()

	err = <-done

	log.Print("Shutting server down")
}
