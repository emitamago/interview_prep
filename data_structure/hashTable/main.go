package main

import (
	"fmt"
)

const ArraySize = 7

// define hashtalbe
type HashTable struct {
	array [ArraySize]*bucket
}

// define bucket
type bucket struct {
	head *bucketNode
}

// define bucketNode
type bucketNode struct {
	key  string
	next *bucketNode
}

// define init for hashtable
func Init() *HashTable {
	result := &HashTable{}
	for i := range result.array {
		result.array[i] = &bucket{}
	}
	return result
}

// define hash function
func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

// define insert for hashtable
func (h *HashTable) Insert(key string) {
	index := hash(key)
	h.array[index].insert(key)
}

// define remove for hastable
func (h *HashTable) Delete(key string) {
	index := hash(key)
	h.array[index].delete(key)
}

// define search for hashtable
func (h *HashTable) Search(key string) bool {
	index := hash(key)
	return h.array[index].search(key)
}

// define insert for bucket
func (b *bucket) insert(k string) {
	if !b.search(k) {
		newHead := &bucketNode{key: k}
		newHead.next = b.head
		b.head = newHead
	} else {
		fmt.Printf("%v is already in the list\n", k)
		return
	}

}

// define remove for bucket
func (b *bucket) delete(key string) {
	if b.head.key == key {
		b.head = b.head.next
		return
	}
	previousToDelete := b.head
	for previousToDelete.next != nil {
		if previousToDelete.next.key == key {
			previousToDelete.next = previousToDelete.next.next
			fmt.Printf("%v is successfully removed\n", key)
			return
		}
		previousToDelete = previousToDelete.next
	}
	fmt.Printf("Could not find %v\n", key)
}

// define search for bucket
func (b *bucket) search(key string) bool {
	currentNode := b.head
	for currentNode != nil {
		if currentNode.key == key {
			return true
		}
		currentNode = currentNode.next
	}
	return false
}

func main() {
	myHashTable := Init()
	list := []string{
		"emi",
		"bond",
		"levi",
		"gio",
		"ram",
	}

	// Inserting key
	for _, v := range list {
		myHashTable.Insert(v)
	}

	// Searching for keys
	answer1 := myHashTable.Search("emi")
	answer2 := myHashTable.Search("Yu")
	fmt.Printf("awnser1 is %v\n", answer1)
	fmt.Printf("awnser2 is %v\n", answer2)

	// Try to insert existing key
	myHashTable.Insert("emi")

	// Delete key
	myHashTable.Delete("emi")
}
