package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

func (q *Queue) Dequeue() int {
   if len(q.items) == 0 { // avoid slice bounds out of range error
        panic("Queue is empty")
   }
   // get first item
   item := q.items[0]
   q.items = q.items[1:] // everything else except the first
   return item
}


// IsEmpty checks if the queue is empty
func (q *Queue) IsEmpty() bool {
    return len(q.items) == 0
}

// Front returns the front element of the queue without removing it
func (q *Queue) Front() int {
    if len(q.items) == 0 {
        panic("Queue is empty")
    }
    return q.items[0]
}

func main(){
  q := Queue{}

  q.Enqueue(1)
  q.Enqueue(2)
  q.Enqueue(3)
  q.Enqueue(4)

  q.Dequeue()

  fmt.Println(q)
}