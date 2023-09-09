# binary tree
- a tree with at most two children
- top most node is the root node
- nodes directly under the root are called children
- nodes at the bottom of the tree are called leaf nodes
- its advantage is speed. In a balanced tree, searching is O(H). Better than O(N)
- using recursion makes it easy to implement
```
  
   |
   |      /O(n)
 t |     /
   |    /____________________________O(log n)
   |   /
   |  /
   |_/_________________________________
          data input
```

- bsts are useful when handling large amouts of data
- balanced tree vs unbalanced tree

              O
            /    \ 
            0     0
          /   \
          0    0

          Balanced tree

          0
            \
             0
              \
               0
         Unbalanced tree

    Binary Tree Example(Balanced - bTree)

                   100
                /        \ 
             52            203
            /  \         /    \ 
          19    76      150    310
        /  \    / \    / \     /  \
       7   24  56 88  135 167 276  423


# insert a new node 56
- start at the root 100, and compare 56 to root
- if 56 < 100, go left (descend to 100's left subtree), 
- if 56 > 100, go right(descend to 100's right subtree)
- descend down the tree and repeat same steps

# find a node with 24 - preorder traversal
- start at the root node
- compare 24 to 100
- if smaller go to the left side, else go to right side
