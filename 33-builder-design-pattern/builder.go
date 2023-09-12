package builder

// build process defines the necessary steps to build a vehicle
// every builder must implement/follow this interface
type BuildProcess interface {
	SetWheels() BuildProcess
	SetSeats() BuildProcess
	SetStructure() BuildProcess

	// retrieve vehicle instance from the builder
	GetVehicle() VehicleProduct
}

// Director is incharge of vehicle(object) construction
type  Director struct {
	builder BuildProcess
}

func (d *Director) Construct(){
	d.builder.SetSeats().SetWheels().SetStructure()
}

// Builders must follow the build process
func (d *Director) SetBuilder(b BuildProcess){
	d.builder = b
}

// the vehicle product, is the vehicle(object) we wish to build
// a vehicle is comprised of wheels, seats and structure
type VehicleProduct struct{
	Wheels, Seats int
	Structure string
}

// the first builder is the car builder
// a builder will store the vehicle product object
type CarBuilder struct {
	v VehicleProduct
}

func (c *CarBuilder) GetVehicle() VehicleProduct {
	return c.v
}

func (c *CarBuilder) SetWheels() BuildProcess {
	c.v.Wheels = 4
    return c
}

func (c *CarBuilder) SetSeats() BuildProcess {
	c.v.Seats = 5
    return c
}

func (c *CarBuilder) SetStructure() BuildProcess {
	c.v.Structure = "Car"
    return c
}

// motorbike structure is the same as the car structure
type BikeBuilder struct {
	v VehicleProduct
}

func (b *BikeBuilder) GetVehicle() VehicleProduct{
	return b.v
}

// define motorbike to have 2 wheels
func (b *BikeBuilder) SetWheels() BuildProcess {
	b.v.Wheels = 2
	return b
}

// define motorbike to have 2 seats
func (b *BikeBuilder) SetSeats() BuildProcess {
	b.v.Seats = 2
	return b
}

func (b *BikeBuilder) SetStructure() BuildProcess {
	b.v.Structure = "Bike"
	return b
}


