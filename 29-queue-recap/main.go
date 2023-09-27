package main

import "fmt"

type Queue struct{
	items []int
}

// we need a pointer receiver
func (s *Queue) Enqueue(i int){
	s.items = append(s.items, i)
}

// Dequeue a value from the Queue and return it
func (s *Queue) Dequeue() int {
   if len(s.items) == 0 {
        panic("Queue is empty")
   }

   item := s.items[0] // get first item
   s.items = s.items[1:] // dequeue first item from queue

   return item
}

func main(){
   myQueue := Queue{}

   myQueue.Enqueue(100)
   myQueue.Enqueue(200)
   myQueue.Enqueue(300)

   myQueue.Dequeue()
   fmt.Println(myQueue)
}