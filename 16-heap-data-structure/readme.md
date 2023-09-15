# heap data structure
- Are an interesting data structure
- Originally introduced as part of an algorithm called heap sort
- Used in implementing priority queues - to take out/extract an item with the highest priority

```
                                 50
                               /     \ 
                              16      48
                             /  \    /   \
                            14   8  34   20
                            / \ / \
                           9  1 5  7

```
- All the levels in the tree are full except for the lowest levels (leaf nodes)
- A binary heap has nodes having at most two children

## max heap
- For a max heap the largest key will be stored in the root node
- Every parent will have a higher key than it s children
- This is useful extract the largest key from the heap.
- Getting the largest key from the heap is fast regardless of the size of heap

## min heap
- Opposite of max heap
- the root node is the smallest key
- each parent wil have a smaller key than its children

```
                                    1
                                 /      \ 
                               8         5
                             /   \      /  \
                            14     9   7   20
                            / \   /  \
                           16 50 34  48

```

## visualization
- heaps can be visualized as a tree but behind the scenes, the keys are stored as an array
- each node on the tree corresponds to an index of that array
- heaps are basically arrays that operate like a tree

```
[0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10]
50  16  48  14  8   34  20  9   1   5   7

```
```
                                 
                                       50[0]
                                    /          \ 
                                16[1]           48[2]
                                /     \         /   \
                            14[3]      8[4]   34[5]  20[6]
                            /   \     /     \
                           9[7] 1[8] 5[9]   7[10]

```
## calculation of index
- you can easily calculate the index of children based on the index of parent

```
Parent Index            Left Child Index
[1]                     [1] * 2 + 1 = [3]

Parent Index            Right Child Index
[1]                     [1] * 2 + 2 = [4]

Parent Index            Left Child Index
[39]                    [39] * 2 + 1 = [79]

```

## obtaining parent index
- left child index is always an odd number
- right child index is an even number

  Parent Index =  [index] - 1 / 2

## Inserting keys in a heap
### Insert and heapify
- Whever we perform an insert, we add a new key to the bottom right of the tree
- The bottom right node will be the last index

```
Suppose we add a new key 63 at index 11

                                 50
                               /     \ 
                              16       48
                             /  \      /   \
                            14   8    34   20
                            / \ / \   /
                            9 1 5  7  63

```
### Heapify Up
- We need to rearrage the nodes(heapify) so that we can maintain the heap property
- We need to keep the parent larger than its children
- We compare the new node having 63 to its parent node having 34
- If the new node is higher (63 > 34), we swap it
- We follow up the tree and keep on repeating the process until it gets to its place

```

                                 63
                               /     \ 
                              16       50
                             /  \      /   \
                            14   8    48   20
                            / \ / \   /
                            9 1 5  7  34

```
## Extracting keys in a heap
### Extract and heapify down
- Heapify process is also needed for taking out an item off the tree
- To extract the highest key in a heap, we can just take out the root
- We need to rearrange the nodes(heapify down) so that we can fill in the empty root while maintaining the heap property

```

                                 63
                               /     \ 
                              16       50
                             /  \      /   \
                            14   8    48   20
                            / \ / \   /
                            9 1 5  7  34

```

### extraction process
- extract the highest key in a heap, 63
- fill in the empty root with the last node 34

```
[0] [1] [2] [3] [4] [5] [6] [7] [8] [9] [10]
34  16  50  14  8   48  20  9   1   5   7

```

- we start the swapping process starting from the top(heapify down)

```
                    34
                  /   \ 
                16    50

```
- 50 becomes the parent by swapping positions with 34

```
                    50
                  /   \ 
                16    34
                     /  \ 
                    48  20
```
- 48 is the larger child, we will swap 34 with 48
```
                    50
                  /   \ 
                16    48
                     /  \ 
                    34  20
```

### Time Complexity(O(h))
- the time complexity for extracting and inserting a node in the heap would depend on the heapify process

## Heapify up
- the number of operations needed to heapify up or down depends on how high the tree is

```

                                 63                 __
                               /     \              |
                              16       50           |
                             /  \     /   \         |h
                            14   8    48   20       |
                            / \ / \   /             |
                            9 1 5  7  34           _|_

```
- we can express insert and extract as O(h), where h is the height of the tree
- if you want to express the time complexity with n, the size of the array, we can replace h with log of n
- the height ans the number of indices have a logarithmic relation
- the time complexity of inserting and extracting from a heap would be O(log n)

## solution 
```
struct MaxHeap
methods -Insert, Extract, maxHeapifyUp, maxHeapifyDown, swap
functions - parent, left, right

```









