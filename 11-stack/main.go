package main

import "fmt"

type Stack struct {
	items []int
}

// stack uses two operations push and pop
func (s *Stack) Push(item int){
	s.items = append(s.items, item)
}

func (s *Stack) Pop() int {

	if len(s.items) == 0 {
        panic("Stack is empty")
    }

	item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1] //returns a slice

	// item, items := s.items[count - 1], s.items[0: count - 1]
	return item
}

// Peek returns the top element from the stack without removing it
func (s *Stack) Peek() int {
    if len(s.items) == 0 {
        panic("Stack is empty")
    }
    return s.items[len(s.items)-1]
}

// IsEmpty checks if the stack is empty
func (s *Stack) IsEmpty() bool {
    return len(s.items) == 0
}


func main(){
	s := Stack{}

	s.Push(1)
	s.Push(2)
	s.Push(3)
	s.Push(4)

	fmt.Println("Peek:", s.Peek()) // Output: Peek: 4
    fmt.Println("Pop:", s.Pop())   // Output: Pop: 4

    fmt.Println("Stack after pop:")
    for !s.IsEmpty() {
        fmt.Println(s.Pop())
    }
}