package main

import (
	"fmt"
	"math/rand"
	"time"
)


func main(){
    ch := make(chan int)
	go process2(ch)

	fmt.Println("Waiting for process")
    
	msg := <- ch  // read message from channel
    fmt.Printf("Process took %d ms\n", msg)
}

func process2(ch chan int){
	n := rand.Intn(3000)
	time.Sleep(time.Duration(n) * time.Millisecond)
    
	ch <- n // send a message to the channel
}