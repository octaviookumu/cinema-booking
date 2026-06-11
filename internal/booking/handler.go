package booking

import (
	"cinema-seat-booking/internal/utils"
	"net/http"
)

type handler struct{
	svc *Service
}

func NewHandler(svc *Service) *handler{
	return &handler{svc}	
}

func (h *handler) ListSeats(w http.ResponseWriter, r *http.Request){
	movieId := r.PathValue("movieId")
	bookings := h.svc.ListBookings(movieId)

	seats := make([]seatInfo, 0, len(bookings))
	for _, b := range bookings {
		seats = append(seats, seatInfo{
			SeatID: b.SeatID,
			UserID: b.UserID,
			Booked: true,
		})
	}
	utils.WriteJSON(w, http.StatusOK, seats)
}

type seatInfo struct {
	SeatID string `json:"seat_id"`
	UserID string `json:"user_id"`
	Booked bool   `json:"booked"`
}