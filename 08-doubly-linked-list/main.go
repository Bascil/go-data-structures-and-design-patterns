package main

import "fmt"

// define a doubly linked list
type List struct {
	head *Node // head points to Node
	tail *Node
}

// node has two parts: data, link
type Node struct {
	data int   // contains the actual data
	next *Node // contains the address(pointer) of the next node in the list
    prev *Node
}

// implement find
func(l *List) Find(value int) *Node{
  return nil
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

// get the previous node
func (n *Node) Prev() *Node {
	return n.prev
}

// add a new node to front of list
func (l *List) AddFront(value int){
	newNode := &Node{data: value, prev: nil, next: nil} // create new node

	if l.head == nil { // if list is empty, assign node to the head
  	  l.head = newNode
	  l.tail = newNode
	  return
	} 
	
	newNode.next = l.head
    l.head.prev = newNode
    l.head = newNode
}

// aDDBack inserts a new node at the back of the list
func (l *List) AddBack(data int) {
    newNode := &Node{data: data, prev: nil, next: nil}

    if l.tail == nil {
        l.head = newNode
        l.tail = newNode
        return
    }

    newNode.prev = l.tail
    l.tail.next = newNode
    l.tail = newNode
}

// Display prints the elements of the doubly linked list
func (l *List) Display() {
    current := l.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
	
    fmt.Println("nil")
}

// ReverseDisplay prints the elements of the doubly linked list in reverse order
func (l *List) ReverseDisplay() {
    current := l.tail
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.prev
    }
    fmt.Println("nil")
}

func main(){
	ll := &List{}

	ll.AddFront(1)
    ll.AddFront(2)
    ll.AddFront(3)
    ll.AddFront(4)

	// display list
    ll.Display()

	// reverse linked list
	ll.ReverseDisplay()
}