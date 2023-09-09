package main

import "fmt"

// To reverse a singly linked list in Go, you can use an iterative approach.

// Define a Node struct and a LinkedList struct.

// Node represents a node in the linked list
type Node struct {
    data int
    next *Node
}

// LinkedList represents the linked list
type LinkedList struct {
    head *Node
}

// InsertFront inserts a new node at the front of the list
func (ll *LinkedList) InsertFront(data int) {
    newNode := &Node{data: data, next: ll.head}
    ll.head = newNode
}

// Display prints the elements of the linked list
func (ll *LinkedList) Display() {
    current := ll.head
    for current != nil {
        fmt.Printf("%d -> ", current.data)
        current = current.next
    }
    fmt.Println("nil")
}

// Reverse reverses the linked list in-place using an iterative approach
func (ll *LinkedList) Reverse() {
    var prev, next *Node
    current := ll.head

    for current != nil {
        next = current.next
        current.next = prev
        prev = current
        current = next
    }

    ll.head = prev
}

func main() {
    ll := LinkedList{}
    
    ll.InsertFront(10)
    ll.InsertFront(5)
    ll.InsertFront(20)
    ll.InsertFront(2)

    fmt.Println("Original Linked List:")
    ll.Display()

    ll.Reverse()

    fmt.Println("Reversed Linked List:")
    ll.Display()
}
