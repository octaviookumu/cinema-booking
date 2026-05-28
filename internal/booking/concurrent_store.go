package booking

import "sync"

type ConcurrentStore struct{
	bookings map[string]Booking // map seat to booking, seat to a booking struct "A2" -> booking
	// Go maps are not concurrency safe
	sync.RWMutex
}

// constructor
func NewConcurrentStore() *ConcurrentStore{
	return &ConcurrentStore{
		bookings: map[string]Booking{},
	}
}

// check whether a seat is already booked
// prevent duplicate booking
// save the booking if the seat is free

// this method is not thread safe, but we'll fix that in the next step
func (s *ConcurrentStore) Book(b Booking) error{
	s.Lock()
	defer s.Unlock()

	if _, exists := s.bookings[b.SeatID]; exists {
		return ErrSeatAlreadyBooked
	}

	s.bookings[b.SeatID] = b
	return nil
}

// list all bookings for a given movie
func (s *ConcurrentStore) ListBookings(movieId string) []Booking{
	s.RLock()
	defer s.RUnlock()

	var result []Booking

	// loop through all bookings and filter by movie ID
	for _, booking := range s.bookings {
		if booking.MovieID == movieId {
			result = append(result, booking)
		}
	}
	return result
}