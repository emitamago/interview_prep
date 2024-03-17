package main

import "fmt"

// create stack type
type Stack struct {
	items []int
}

// create push
func (s *Stack) push(i int) {
	s.items = append(s.items, i)
}

// create pop
func (s *Stack) pop() (i int) {
	toRemove := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return toRemove
}

func main() {
	myStack := Stack{}
	myStack.push(6)
	myStack.push(20)
	myStack.push(34)
	fmt.Printf("my stack is %v\n", myStack)
	myStack.pop()
	fmt.Printf("my stack is %v\n", myStack)
}
