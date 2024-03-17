package main

import (
	"fmt"
)

const AlphabetNumber = 26

// Define node
type Node struct {
	children [AlphabetNumber]*Node
	IsEnd    bool
}

// Define Trie
type Trie struct {
	root *Node
}

// Initiate trie
func InitTrie() *Trie {
	result := &Trie{&Node{}}
	return result
}

// Define insert
func (t *Trie) Insert(w string) {
	currentNode := t.root
	for _, v := range w {
		charIndex := v - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.IsEnd = true
}

// Define search
func (t *Trie) Search(w string) bool {
	currentNode := t.root
	for _, v := range w {
		charIndex := v - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.IsEnd {
		return true
	}
	return false
}

func main() {
	testTrie := InitTrie()
	names := []string{
		"emi",
		"tsukuda",
		"gio",
		"ram",
		"levi",
		"bonbon",
	}

	for _, v := range names {
		testTrie.Insert(v)
	}

	a1 := testTrie.Search("emi")
	a2 := testTrie.Search("levi")
	a3 := testTrie.Search("tom")

	fmt.Printf("answers are %v, %v, %v\n", a1, a2, a3)
}
