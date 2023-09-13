package abstract_factory

type SportMotorbike struct{}

func (s *SportMotorbike) GetWheels() int {
	return 5
}

func (s *SportMotorbike) GetSeats() int {
	return 1
}