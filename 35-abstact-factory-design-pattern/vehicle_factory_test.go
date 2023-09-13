package abstract_factory

import "testing"

func TestMotorbikeFactory(t *testing.T){
	motorbikeF, err := GetVehicleFactory(MotorbikeFactoryType)

	if err != nil {
		t.Fatal(err)
	}

	motorbikeVehicle, err := motorbikeF.GetVehicle(SportMotorbikeType)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Motorbike vehicle has %d wheels and %d seats", motorbikeVehicle.GetWheels(), motorbikeVehicle.GetSeats())
	sportBike, ok := motorbikeVehicle.(Motorbike)

	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Sportbike has type %d\n",sportBike.GetType())
}

func TestCarFactory(t *testing.T){
	carF, err := GetVehicleFactory(3)

	if err != nil {
		t.Fatal("Car factory with id 3 not recognized")
	}

	carF, err = GetVehicleFactory(CarFactoryType)

	if err != nil {
		t.Fatal(err)
	}

	carVehicle, err := carF.GetVehicle(LuxuryCarType)

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Car vehicle has %d wheels and %d seats", carVehicle.GetWheels(), carVehicle.GetSeats())
	luxuryCar, ok := carVehicle.(Car)

	if !ok {
		t.Fatal("Struct assertion has failed")
	}

	t.Logf("Luxury car has %d doors\n",luxuryCar.GetDoors())
}