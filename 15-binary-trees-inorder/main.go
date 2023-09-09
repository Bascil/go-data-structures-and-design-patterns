package main

import "fmt"

// TreeNode represents a node in the binary tree
type TreeNode struct {
    data  int
    left  *TreeNode
    right *TreeNode
}

// BinaryTree represents a binary tree
type BinaryTree struct {
    root *TreeNode
}

// Insert inserts a new node into the binary tree
func (bt *BinaryTree) Insert(data int) {
    newNode := &TreeNode{data: data, left: nil, right: nil}

    if bt.root == nil {
        bt.root = newNode
        return
    }

    queue := []*TreeNode{bt.root}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        if node.left == nil {
            node.left = newNode
            return
        } else {
            queue = append(queue, node.left)
        }

        if node.right == nil {
            node.right = newNode
            return
        } else {
            queue = append(queue, node.right)
        }
    }
}

// InorderTraversal performs an inorder traversal of the binary tree
func (bt *BinaryTree) InorderTraversal(node *TreeNode) {
    if node != nil {
        bt.InorderTraversal(node.left)
        fmt.Printf("%d\n", node.data)
        bt.InorderTraversal(node.right)
    }
}

func main() {
    bt := &BinaryTree{}
    
    bt.Insert(10)
    bt.Insert(5)
    bt.Insert(20)
    bt.Insert(2)
    bt.Insert(7)

    fmt.Println("Inorder Traversal:")
    bt.InorderTraversal(bt.root)
}
