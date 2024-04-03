package main

import "fmt"

// define a node
type Node struct {
	Key   int
	Right *Node
	Left  *Node
}

// define Insert
func (n *Node) Insert(key int) {
	if n.Key < key {
		if n.Right == nil {
			n.Right = &Node{Key: key}
		} else {
			n.Right.Insert(key)
		}
	} else if n.Key > key {
		if n.Left == nil {
			n.Left = &Node{Key: key}
		} else {
			n.Left.Insert(key)
		}
	}
}

// define Search
func (n *Node) Search(key int) bool {
	// At the end of tree
	if n == nil {
		return false
	}

	if n.Key < key {
		return n.Right.Search(key)
	} else if n.Key > key {
		return n.Left.Search(key)
	}

	// key is not larger than or smaller than right child or left child == match
	return true
}

func main() {
	tree := &Node{Key: 100}
	list := []int{200, 50, 23, 7, 98}
	for _, v := range list {
		tree.Insert(v)
	}
	answer1 := tree.Search(50)
	answer2 := tree.Search(10000)
	fmt.Printf("answer1 is %v\n", answer1)
	fmt.Printf("answer2 is %v\n", answer2)
}
