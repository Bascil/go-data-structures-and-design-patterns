package main

import "fmt"

type FooOpts struct {
    // optional arguments
    Value string
}

func main(){

  nums := []int32{1,2,3}

  foo()
  foo(1)
  foo(1,2,3)

  fmt.Printf("%d\n", sum(1,2,4))
  fmt.Printf("%d\n", sum(1,2,4, 5))
  fmt.Printf("%d\n", sum2(nums))

  NewFooWithOpts("Make it work please", &FooOpts{Value: " World"})
}

// use variadic args.
func foo(params ...int) {
    fmt.Println(len(params))
}

//you can name the return value in the return type
func sum(args ...int) (result int) { 
	for _, v := range args {
      result += v
	}

	return result // or simply return
}

func sum2(args []int32) int32 {
	sum := int32(0)
	for _, v := range args{
       sum += v
	}

	return sum
}


func NewFooWithOpts(mandatory string, opts *FooOpts) {
    if (&opts) != nil {
        fmt.Println("Hello " + opts.Value)
    } else {
        fmt.Println("Hello")
    }
}