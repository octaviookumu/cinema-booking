package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /movies", listMovies)

	mux.Handle("GET /", http.FileServer(http.Dir("static")))

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}

var movies = []movieResponse{
	{ID: "inception", Title: "Inception", Rows: 5, SeatsPerRow: 8},
	{ID: "dune", Title: "Dune", Rows: 4, SeatsPerRow: 6},
}

// handler for listing movies
func listMovies(w http.ResponseWriter, r *http.Request){
	WriteJSON(w, http.StatusOK, movies)
}

// helper function to write JSON responses
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(v)
}

// what the frontend is expecting
type movieResponse struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Rows int `json:"rows"`
	SeatsPerRow int `json:"seats_per_row"`
}