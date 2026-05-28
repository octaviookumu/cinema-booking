package booking

type Service struct {
	// dependency injection of the store
	store BookingStore
}

func NewService(store BookingStore) *Service {
	return &Service{store}
}

func (s *Service) Book(b Booking) error {
	return s.store.Book(b)
}