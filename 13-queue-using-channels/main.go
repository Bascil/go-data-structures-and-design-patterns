package main

import "fmt"

type Queue struct {
	items chan int
}

func (q *Queue) Enqueue(i int){
	q.items <- i // push an item onto  the channel
}

func (q *Queue) Dequeue() int {
	return <- q.items // read from channel
}

func main(){
   queue := Queue{
	items: make(chan int, 16),
   }

   queue.Enqueue(1)
   queue.Enqueue(2)
   queue.Enqueue(3)

   // drain the channel
   fmt.Println(queue.Dequeue())
   fmt.Println(queue.Dequeue())
}