## builder design pattern
- allows you to reuse the same algorithm to create many diffent types of objects
- commonly described as the relationship between the director, a few builders and the product they build
- director is incharge of vehicle(object) construction but builders are the ones that build the actual vehicle
- the director entity in the builder pattern is a clear candidate for singleton pattern
- the process of creating a vehicle (object) is the same for every kind of vehicle.
- you would use more or less the same technique(build process) to build a car as you would build a motorbike or a bus. 
- the difference may just be the number of seats and size
- you select the type of vehicle, assemble the structure, place the wheels and place the seats
- this pattern helps maintain an unpredictable number of products/objects using a common construction algorithm (build process) used by the director
- object creation is separated from object user. You can abstract so that the user does not need to know the details of object creation(build process)

## build process
- we use a common build process to build different kinds of vehicles
- build process interface specifies what one must comply to be part od the process

## example
- we take an example of vehicle manufacturing to implement this design pattern

## algorithm
1. You must have a manufacturing type e.g ManufacturingDirectorType
2. Car VehicleProduct will have 4 wheets, 5 seats and a car structure
3. Motorbike Vehicle product will have 2 wheels, 2 seats and a motorbike structure
4. Vehicle product must be open for modifications

## run tests
```
go test -v -run=TestBuilder
```

