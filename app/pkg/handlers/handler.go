package handlers

import (
	"fmt"
	"net/http"
	"prw_server/app/pkg/api"
)

type Handler struct {
	JokeClient api.Client
}

func NewHandler(jokeClient api.Client) *Handler {
	return &Handler{
		JokeClient: jokeClient,
	}
}

func (h *Handler) Hello(w http.ResponseWriter, r *http.Request) {
	joke, err := h.JokeClient.GetJoke()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Fprint(w, joke.Joke)
}
