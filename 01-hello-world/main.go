package main // define main package

// import dependencies
import "fmt"

func hello(name string){
	fmt.Println("Hello", name)
}

func main() {
	hello("Basil")
	hello("Gophers")
}