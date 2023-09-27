package main

import "fmt"

var count int

// define a node struct, each node points to its left and right child
type Node struct{
	key int
	
	Left *Node
	Right *Node
}

// we use a pointer receiver
func (n *Node) Insert(k int){
	// if k is larger than the key
	if k > n.key { // move right
		if n.Right == nil { //if child is empty, create a new node
			n.Right = &Node{key: k}
		} else {
			// keep checking(descend down the tree), recursively call the insert
			n.Right.Insert(k)
		}
	} else { // move left
		if n.Right == nil {
			n.Left = &Node{key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

func (n *Node) Search(k int) bool {
	count++
	if n == nil { // if node is nil, its end of tree, return false
		return false
	}

	if n.key == k {
		return true
	}
	// ensure you are not traversing the whole tree
	fmt.Printf("%d\n", n.key)

	if k > n.key { //move right
      return  n.Right.Search(k) // descend down the tree, keep searching till you find a match
	} else {
      return n.Left.Search(k)
	}
}

func main(){
   tree := &Node{key: 100}

   tree.Insert(52)
   tree.Insert(203)
   tree.Insert(19)
   tree.Insert(76)
   tree.Insert(150)
   tree.Insert(310)
   tree.Insert(7)
   tree.Insert(24)
   tree.Insert(88)
   tree.Insert(276)

   fmt.Println(tree.Search(88))
   fmt.Println(count)
}