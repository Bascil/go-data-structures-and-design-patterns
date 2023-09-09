package main

// define a singly linked list

type List struct {
	head *Node // head points to Node
	tail *Node
}

// node has two parts: data, link
type Node struct {
	data int   // contains the actual data
	next *Node // contains the address(pointer) of the next node in the list
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

// add a new node to list
func (l *List) Add(value int){
	node := &Node{data: value} // create new node

	if l.head == nil { // if list is empty, assign node to the head
  	  l.head = node
	} else {
		// if the list has some element already
		// previous tail.next will be the node
		l.tail.next = node 
	}

	l.tail = node // in every case assign it to the tail as well
}

func main(){
	l := &List{}
	l.Add(1)
	l.Add(2)
	l.Add(3)

	n := l.First()
	// println(n.data)

	// n = n.Next()
	// println(n.data)

	// last element.next is nil
	// iterate over elements till we find a nil element
	for {
		println(n.data)
		n = n.Next()
		if n == nil {
			break
		}
	}
}