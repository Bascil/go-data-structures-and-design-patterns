package abstract_factory

type LuxuryCar struct{}

func(f *LuxuryCar) GetDoors() int {
	return 4
}

func(f *LuxuryCar) GetWheels() int {
	return 4
}

func(f *LuxuryCar) GetSeats() int {
	return 5
}