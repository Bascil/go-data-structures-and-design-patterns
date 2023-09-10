package main

import "fmt"

type Stack struct{
	items []int
}

// we need a pointer receiver
func (s *Stack) Push(i int){
	s.items = append(s.items, i)
}

// pop a value from the stack and return it
func (s *Stack) Pop() int {
   if len(s.items) == 0 {
        panic("Stack is empty")
   }

   last := len(s.items) - 1
   item := s.items[last]
   s.items = s.items[:last] // 0:last // returns a slice

   return item
}

// return the top element from the stack without removing it
func (s *Stack) Peek() int {
    if len(s.items) == 0 {
        panic("Stack is empty")
    }
    return s.items[len(s.items)-1]
}

func main(){
   myStack := Stack{}

   myStack.Push(100)
   myStack.Push(200)
   myStack.Push(300)

   myStack.Pop()
   fmt.Println(myStack)
}