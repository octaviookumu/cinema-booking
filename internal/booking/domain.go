package booking

import "errors"

var(
	ErrSeatAlreadyBooked = errors.New("seat is already taken")
)

// Booking represents a confirmed seat reservation.
type Booking struct {
	ID      string
	MovieID string
	SeatID  string
	UserID  string
	Status  string
}

// the interface we'll use dependency injection for
type BookingStore interface {
	Book(b Booking) error
	ListBookings(movieID string) []Booking
}