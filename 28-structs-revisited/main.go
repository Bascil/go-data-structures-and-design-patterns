package main

import "fmt"

type Person struct {
	Name, Surname string
	Hobbies []string
	id string // private
}

func (p *Person) GetFullName() string {
	return fmt.Sprintf("%s %s", p.Name, p.Surname)
}

func main(){
	person := Person{
		Name: "Basil",
		Surname: "Ndonga",
		Hobbies: []string{"hicking", "coding"},
		id: "12345-89",
	}

	fmt.Printf("%s likes %s\n", person.GetFullName(), person.Hobbies[0])
}