package builder

import "testing"

func TestBuilder(t *testing.T){
	// create a build process for every vehicle
	buildProcess := Director{}

	// create a car builder
	carBuilder := &CarBuilder{}

	// pass car builder to process using the setBuilder() method
	buildProcess.SetBuilder(carBuilder)

	// since the director now knows what to construct, we call the construct method
	buildProcess.Construct()

	// once we have all the pieces for our car, we call the get vehicle method 
	// to return a car instance
	car := carBuilder.GetVehicle()

	// check if the car has 4 wheels
	if car.Wheels != 4 {
    	t.Errorf("wheels of a car must be 4 and we have %d\n",car.Wheels)
	}

	// check if the car has 5 seats
	if car.Seats != 5 {
		t.Errorf("seats of a car must be 5 and we have %d\n",car.Seats)
	}

	// check if structure is 'Car'
	if car.Structure != "Car" {
		t.Errorf("structure of a car must be 'Car' and we have %s\n", car.Structure)
	}

	// create a bike builder
	bikeBuilder := &BikeBuilder{}

	// pass car builder to process using the setBuilder() method
	buildProcess.SetBuilder(bikeBuilder)

	// since the director now knows what to construct, we call the construct method
	buildProcess.Construct()

	bike := bikeBuilder.GetVehicle()
	//bike.Seats = 1

	// check if the car has 5 seats
	if bike.Seats != 2 {
			t.Errorf("seats of a bike must be 2 and we have %d\n",bike.Seats)
	}

	if bike.Wheels != 2 {
		t.Errorf("wheels of a bike must be 2 and we have %d\n",car.Wheels)
	}

	// check if structure is 'Bike'
	if bike.Structure != "Bike" {
		t.Errorf("structure of a bike must be 'Bike' and we have %s\n", bike.Structure)
	}
}