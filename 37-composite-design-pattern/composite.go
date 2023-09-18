package composite

type Athlete struct{}

func (a *Athlete) Train() {
	println("Training")
}

func Swim(){
	println("Swiming")
}

// create a composite swimmer that has athlete struct inside int
type CompositeSwimmerA struct{
	MyAthlete Athlete // has MyAthlete field of type Athlete

	// in go functions can be used as paramenters or fields just like any variable - pointer to a function
	MySwim *func() // stores a func type(closure)
}

type Trainer interface {
	Train()
}

type Swimmer interface{
	Swim()
}

type SwimmerImplementor struct{}

func (s *SwimmerImplementor) Swim(){
	println("swimming!")
}

type CompositeSwimmerB struct{
	Trainer
	Swimmer
}

type Animal struct{}

func(r *Animal) Eat(){
	println("eating!")
}

// shark object embeds the animal object
// you can embed objects within object to look alot like inheritance
type Shark struct{
	Animal
	Swim func()
}