package main

import "fmt"

// size of hash table array
const ArraySize = 7

// hash table struct holds an array
type HashTable struct{
   array [ArraySize]*bucket
}

// bucket is a linked list in each slot of the array
type bucket struct {
	// put the address of the head node in the bucket
	head *bucketNode 
}

type bucketNode struct {
	key string
	next *bucketNode
}

// initialize the hash table
func Init() *HashTable{
	// create a new hash table
	result := &HashTable{}

	// each slot of the array has to hold a bucket
	for i := range result.array{
		result.array[i] = &bucket{}
	}

	// return the address of the hash table
	return result
}

// hash function  computes the hash code
func hash(key string) int {
	//set the ascii code for each character 
	// sum it up and divide by array size
	sum := 0
	for _, v := range key{
		sum += int(v)
	}

	return sum % ArraySize
}

// insert will take in a key and add it to HT array
func(h *HashTable) Insert(key string){
  index := hash(key)

  // insert the node in the bucket
  h.array[index].insert(key)
}

// search takes in a key and return true if 
// that key is stored in the hash table
func(h *HashTable) Search(key string) bool {
	index := hash(key) // get the hash code
	return h.array[index].search(key) // bool
}

// delete will take in a key and delete it from the HT
func(h *HashTable) Delete(key string){
	index := hash(key) // get the hash code
	h.array[index].delete(key)
}

//----------------------------------------//
// bucket methods
func (b *bucket) insert(k string){

	if !b.search(k) {
		// get a key and create a bucket node for that key
		newNode := &bucketNode{key:k }

		// next node becomes the current head
		newNode.next = b.head

		// head will become the new node
		b.head = newNode
	} else {
		fmt.Println(k, "Already exists")
	}
	
}

// seach takes in a key and returns
// true if the bucket has that key
func(b *bucket) search(k string) bool {
   currentNode := b.head

   // keep looping till be find a match
   for currentNode != nil {
      if currentNode.key == k {
		return true
	  }

	  // update the node to next node
	  currentNode = currentNode.next
   }

   return false
}

// delete will take in a key and delete 
// the node from the bucket
// similar to search except that we use previous node
func (b *bucket) delete(k string){
// if key of head is the match to be deleted
// if the matching key is head
   if b.head.key == k {
	b.head = b.head.next // reset the head of bucketto second node
	return
   }

   previousNode := b.head 
   for previousNode.next != nil {
	// loop through and update
	if previousNode.key == k {
		previousNode.next = previousNode.next.next
	}
	previousNode = previousNode.next
   }
}

func main(){
	hashTable := Init()
	list := []string{"ERIC", "KENNY", "KYLE", "STAN", "RANDY", "BUTTERS", "TOKEN"}

	for _, v := range list{
		hashTable.Insert(v)
	}

	hashTable.Delete("STAN")
	fmt.Println(hashTable.Search("KENNY"))
	fmt.Println(hashTable.Search("STAN"))
}