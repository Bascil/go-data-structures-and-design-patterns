package main

import "fmt"

// Node represents a node in the doubly linked list
type Node struct {
    data int
    prev *Node
    next *Node
}

// DoublyLinkedList represents the doubly linked list
type DoublyLinkedList struct {
    head *Node
    tail *Node
}

// InsertFront inserts a new node at the front of the list
func (dll *DoublyLinkedList) InsertFront(data int) {
    newNode := &Node{data: data, prev: nil, next: nil}

    if dll.head == nil {
        dll.head = newNode
        dll.tail = newNode
        return
    }

    newNode.next = dll.head
    dll.head.prev = newNode
    dll.head = newNode
}

// InsertBack inserts a new node at the back of the list
func (dll *DoublyLinkedList) InsertBack(data int) {
    newNode := &Node{data: data, prev: nil, next: nil}

    if dll.tail == nil {
        dll.head = newNode
        dll.tail = newNode
        return
    }

    newNode.prev = dll.tail
    dll.tail.next = newNode
    dll.tail = newNode
}

// Display prints the elements of the doubly linked list
func (dll *DoublyLinkedList) Display() {
    current := dll.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}

// ReverseDisplay prints the elements of the doubly linked list in reverse order
func (dll *DoublyLinkedList) ReverseDisplay() {
    current := dll.tail
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.prev
    }
    fmt.Println("nil")
}

func main() {
    dll := DoublyLinkedList{}
    
    dll.InsertBack(10)
    dll.InsertFront(5)
    dll.InsertBack(20)
    dll.InsertFront(2)

    fmt.Println("Doubly Linked List:")
    dll.Display()

    fmt.Println("Reversed Doubly Linked List:")
    dll.ReverseDisplay()
}
