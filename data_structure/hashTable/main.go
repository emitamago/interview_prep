package main

import "fmt"

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
// func (h *HashTable) Delete (key string) {
// 	index := hash(key)
// }
// define search for hashtable
// func (h *HashTable) Search (key string) bool {
// 	index := hash(key)
// }

// define insert for bucket
func (b *bucket) insert(k string) {
	newHead := &bucketNode{key: k}
	newHead.next = b.head
	b.head = newHead
}

// define remove for bucket

// define search for bucket

func main() {
	myHashTable := Init()
	fmt.Printf("my hashtable is %v\n", myHashTable)
	testBucket := &bucket{}
	testBucket.insert("emi")
	fmt.Printf("my bucket is %v\n", testBucket)
}
