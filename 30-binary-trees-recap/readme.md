# binary tree
- all the data structures discussed have been linear data structures
- Binary tree are part of tree structures
- a tree with at most two children
- each part of the tree is called a node
- each connection between nodes is called edge
- top most node is the root node
- nodes directly under the root are called children
- nodes of the same parent are called sibling nodes
- nodes at the bottom of the tree are called leaf nodes
- leaf nodes have no children
- a subtree is part of a tree which in itself is a tree 52,19,76 form a subtree
- 52 and 19 are ancestors to node 7
- each node is greater than every node in its left subtree

## advantages
### speed
- you can locate data in a large tree with very few comparisons
- In a balanced tree, searching is O(H). Better than O(N)

### recursion
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

              O  root
            /    \ 
            0     0 parent
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
- if 56 <= 100, go left (descend to 100's left subtree), 
- if 56 > 100, go right(descend to 100's right subtree)
- descend down the tree and repeat same steps

# find a node with 24 - preorder traversal
- start at the root node
- compare 24 to 100
- if smaller go to the left side, else go to right side

## binary tree applications
- searching algorithms
- hierarchical data representations in file systems, org charts, html or xml structures
- databases - for efficient storage and retrieval of data allowing lookup, insertion, deletion
- balanced trees(btrees) - used in databases to maintain a balanced structure ensuring efficient performance of databases such as in database indexing. Indexng speeds up searching and sorting of records in databases.
- filesystems - file systems are often represented as trees allowing for effiecient navigation of files and directories
- decision trees in Machine Learning - used in classification algorithms.
- game trees - used to model possible moves and outcomes in games like chess and tic tac toe
- routing tables in networking - to store routing information in networking devices allowing efficient routing of data packets
- XML/HTML DOM - to represent hierarchical structure of documents in XML and HTML. This allows efficient manipulation and traversing(navigation) of the document
- binary heaps for priority queues - Binary heaps are a special type of binary tree to implement priority queues. It is a fundamental data structure for algorithms like Dijkstas algorithm.

## btrees
A B-tree is a self-balancing tree data structure that maintains sorted data and allows searches, sequential access, insertions, and deletions in logarithmic time. The B-tree generalizes the binary search tree, allowing for nodes with more than two children.[2] 

![B-trees](btree.png)

Each node in a B-Tree can contain multiple keys, which allows the tree to have a larger branching factor and thus a shallower height. This shallow height leads to less disk I/O, which results in faster search and insertion operations. B-Trees are particularly well suited for storage systems that have slow, bulky data access such as hard drives, flash memory, and CD-ROMs.

Unlike other self-balancing binary search trees, the B-tree is well suited for storage systems that read and write relatively large blocks of data, such as databases and file systems.

[More](https://www.geeksforgeeks.org/introduction-of-b-tree-2/)







