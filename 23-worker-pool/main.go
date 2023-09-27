package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// create input and out channel
	in := make(chan int)
	out := make(chan int)

	for i := 0; i < 500; i++ {
		go worker(in, out)
	}

	go producer(in)

	for n := range out {
		fmt.Printf("<- Recv job: %d\n", n)
	}
}

func worker(in, out chan int) {
	for {
		n := <- in
		time.Sleep(time.Duration(rand.Intn(3000))* time.Millisecond)

		out <- n
	}
}

func producer(ch chan<- int) {
	i := 0
	for {
		fmt.Printf("-> Send job: %d\n",i )
		ch <- i
		i++
	}
}