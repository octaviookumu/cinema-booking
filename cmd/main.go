package main

import (
	"cinema-seat-booking/internal/booking"
	"log"
	"net/http"

	"cinema-seat-booking/internal/adapters/redis"
	"cinema-seat-booking/internal/utils"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /movies", listMovies)

	mux.Handle("GET /", http.FileServer(http.Dir("static")))

	store := booking.NewRedisStore(redis.NewClient("localhost:6379"))
	svc := booking.NewService(store)

	bookingHandler := booking.NewHandler(svc)

	mux.HandleFunc("GET /movies/{movieId}/seats", bookingHandler.ListSeats)

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
	utils.WriteJSON(w, http.StatusOK, movies)
}


// what the frontend is expecting
type movieResponse struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Rows int `json:"rows"`
	SeatsPerRow int `json:"seats_per_row"`
}