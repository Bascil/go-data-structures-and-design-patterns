package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleep(){
  time.Sleep(time.Duration(rand.Intn(3000))* time.Millisecond)
}

// create a write only channel
func producer(ch chan<- int){
	for {
		sleep() // avoid blasting the consumer with messages
		n := rand.Intn(100) // generate a random number from 0 to 100
		fmt.Printf(" -> %d\n", n)
		ch <- n
	}
}

// create a read only channel
func consumer(ch <-chan int, name string ){
  for n := range ch {
	 fmt.Printf("Consumer %s <- %d\n", name, n)
  }
}

func fanOut(chA <-chan int, chB, chC chan<-int){
	// read from chan A
	for n := range chA{
		if n < 50  { //send lowe half to channel B
			chB <- n
		} else {
			chC <-n // otherwise send to chan C
		}
	}
}

func main(){
  chA := make(chan int)
  chB := make(chan int)
  chC := make(chan int)

  // producer will not be named as we have a single producer
  go producer(chA)

  // create a consumer to read from channel
  go consumer(chB, "B")
  go consumer(chC, "C")

  fanOut(chA, chB, chC)
}