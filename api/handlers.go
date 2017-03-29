package main

import (
	"encoding/json"
	"log"
	"net/http"
	"github.com/rebecabordini/movies-api/movie"
)

const ADDR = "127.0.0.1:8081"

type MovieHandler struct{}

func (h *MovieHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	encoder := json.NewEncoder(w)
	movie := movie.Person{ Name: "Sofia Copolla", Age: 42 }
	err := encoder.Encode(movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func main() {
	handler := &MovieHandler{}
	log.Printf("Running server on ", ADDR)
	log.Fatal(http.ListenAndServe(ADDR, handler))
}