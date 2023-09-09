package main

func hello(){
	println("Hello")
}

func sum(x, y int)int {
	return x + y
}

func swap(x, y int) (int, int) {
	return y, x
}

func main(){
 // hello()
 // println(sum(10, 29))
//  a, b := swap(10, 20)
//  println(a)
//  println(b)

// using closures
// a closure is a function  assigned to a variable
sum := func(x, y int) int {
	return x + y
}
println(sum(4,6))

}