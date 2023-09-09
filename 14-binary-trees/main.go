package main

type Tree struct{
	node *Node
}

type Node struct{
	value int
	left *Node
	right *Node
}

func (t *Tree) Insert(value int) *Tree {
   if t.node == nil {
	t.node = &Node{value: value}
   } else {
	t.node.Insert(value) // node has an insert function
   }
   return t
}

func (n *Node) Insert(value int){
	if value <= n.value {
		if n.left == nil {
			n.left = &Node{value: value}
		} else {
			n.left.Insert(value)
		}
	} else {
		if n.right == nil {
			n.right = &Node{value:value}
		} else {
			n.right.Insert(value)
		}
	}
}

func (n *Node) Exists(value int) bool {
	if n == nil {
		return false
	} 

	if n.value == value {
		return true
	}

	// ensure you are not traversing the whole tree
	println("Value:", n.value)

	if value <= n.value {
		return n.left.Exists(value)
	} else {
		return n.right.Exists(value)
	}
}

func Display(n *Node){
	if n == nil {
		return 
	}

	println(n.value)

	Display(n.left)
	Display(n.right)
}

func main(){
   tree := Tree{}
   tree.Insert(10).Insert(8).Insert(20).Insert(9).Insert(0).Insert(15).Insert(25)

   Display(tree.node)
   println(tree.node.Exists(25))
}