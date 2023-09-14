package main

import "fmt"

type MaxHeap struct{
	items []int
}

// add an element to the heap
func (h *MaxHeap) Insert(key int){
	h.items = append(h.items,key)

	// rearrange to maintain heap property
	h.maxHeapifyUp(len(h.items) - 1)
}

// extract(remove) the largest key from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.items[0]

	if len(h.items) == 0 {
		fmt.Println("unable to extract because array length is 0")
		return -1 // get out of method
	}

	l := len(h.items) - 1

	// get the last index and set it to the root
	h.items[0] = h.items[l] // set last index to the first index

	// remove the last key
	h.items = h.items[:l]

	h.maxHeapifyDown(0)

	return extracted
}

// heapify from botton to top
func (h *MaxHeap) maxHeapifyUp(index int){
 // swap when the current index is larger than its index

 for h.items[index] > h.items[parent(index)] {
   // if key is larger we swap
   h.Swap(parent(index), index) // swap till it finds its place

   // update the loop - set the new parent
   index = parent(index)
 }
}

// get the last child, compare it to current index
func (h *MaxHeap) maxHeapifyDown(index int){
 // get the larger child, compare with the current index and then swap it
 // repeat till index has no children
 // if left child index is smaller than last index of the array, then the index
 // has at least one child
 lastIndex := len(h.items) - 1

 l, r := left(index), right(index)

 // define childToCompare
 childToCompare := 0

 // loop while index has at least one child
for l <= lastIndex {
	// if the left is large, we assign the left child to child to compare
	if l == lastIndex {
		childToCompare = l
	} else if h.items[l] > h.items[r] { // left key is larger than right key
      childToCompare = l
	} else { // when right key is larger
      childToCompare = r
	}

	// compare array value of current index to larger child and swap if swaller
	if h.items[childToCompare] > h.items[index] {
		h.Swap(index, childToCompare)

		// update variables in the loop
		l, r = left(index), right(index)
	} else {
		// if value of current index is larger than the larger child, it has found its place
		return 
	}
  }
}

func (h *MaxHeap) Swap(index1, index2 int){
	h.items[index1], h.items[index2] = h.items[index2], h.items[index1]
}

func parent(index int) int{
	return (index - 1)/2
}

func left(index int) int {
	return (2 * index) + 1
}

func right(index int) int {
	return (2 * index) + 2
}

func main(){
	m := &MaxHeap{}
	buildHeap := []int{10,20,30,5,7,9,11,13,15,17}
	for _, val := range buildHeap {
		m.Insert(val)
		fmt.Println(m)
	}

	
	for i := 0; i < 5; i++ { // extract 5 times
		m.Extract()
		fmt.Println(m)
	}
}