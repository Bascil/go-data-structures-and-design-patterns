package main

import "fmt"

type Node struct {
	data int
	next *Node
	prev *Node
}

type List struct {
	head *Node
	tail *Node
}

// get the first node
func (l *List) First() *Node {
	return l.head
}

// get the last node
func (l *List) Last() *Node {
	return l.tail
}

// get the next node
func (n *Node) Next() *Node {
	return n.next
}

// we need a pointer receiver
func (l *List) Insert(data int){
	newNode := &Node{data: data, prev: nil, next: nil}

    if l.head == nil {
        l.head = newNode
        l.tail = newNode
        return
    }

    newNode.next = l.head
    l.head.prev = newNode
    l.head = newNode
}

// Display prints the elements of the  linked list
func (l *List) Display() {
    current := l.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}

// ReverseDisplay prints the elements of the doubly linked list in reverse order
func (l *List) Reverse() {
    current := l.tail
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.prev
    }
    fmt.Println("nil")
}

func main(){
  myList := List{}

  myList.Insert(51)
  myList.Insert(31)
  myList.Insert(73)
  myList.Insert(90)

  myList.Display()
  myList.Reverse()
}