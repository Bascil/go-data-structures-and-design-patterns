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
func producer(ch chan<- int, name string){
	for {
		sleep()
		n := rand.Intn(100) // generate a random number from 0 to 100
		fmt.Printf("Producing %s -> %d\n", name, n)
		ch <- n
	}
}

// create a read only channel
func consumer(ch <-chan int ){
  for n := range ch {
	 fmt.Printf("Consuming <- %d\n", n)
  }
}

// chA, chB are read only channels, chC is write only
func fanIn(chA, chB <-chan int, chC chan<-int){
	var n int //hold value from any  channel

	for {
		//read from chan A or B, send message to chan C
		select {
		case n = <- chA:
			chC <- n

		case n = <-chB:
			chC <- n
		}
	}
}


func main(){
  chA := make(chan int)
  chB := make(chan int)
  chC := make(chan int)

  // to producers sending messages to channel A and B
  go producer(chA, "A")
  go producer(chB, "B")

  // create a consumer to read from channel
  go consumer(chC)

  //fanIn receives input from chan A and C, sends output to chan C
  fanIn(chA, chB, chC)

}