# composite design pattern
- it is a structural design pattern
- helps shape applications with commonly used structures and relationships
- Go encourages use of composition due to its lack of inheritance
- Go favours composition over inheritance.

## why composite
- we create hierarchies and trees of objects
- objects have different objects with their own fields and methods inside them
- this approach is very powerful and solves many problems of inheritance

### typical inheritance problem
- Athlete Class has a Train() method
- Swimmer Class has a Swim() method
- Swimmer Class inherits Train() method from the Athlete Class
- Cyclist who is also an Athlete declares a Ride() method
- Imagine an Animal Class that Eats() and also Barks()

### solution
- make a Swimmer interface with the Swim method and make the Swimmer, Athlete and Fish implement it.

- the objective is to avoid hierarchy hell where the application complexity could grow and code clarity be affected

- we solve this problem using an `idiomatic` approach. In Go we can use two types of composition ie direct composition and embedding composition

#### direct composition
- we shall make use of direct composition
- is a pure structural design pattern
- We shall have an Athlete, Swimmer, Animal and Fish structs
- The Athlete struct will have Train() method
- We shall have an Animal Struct with Eat() method, animal must eat
- We shall have a Fish Struct with Swim() method
- Swimmer and Fish structs must share the code
- Swim() method is shared with Swimmer struct

#### Note
The zero initialization feature is implemented at the go compiler