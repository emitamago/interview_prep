package main

import "fmt"

// Define MaxHeap
type MaxHeap struct {
	array []int
}

// Define insert
func (m *MaxHeap) Insert(key int) {
	m.array = append(m.array, key)
	m.maxHeapifyUp(len(m.array) - 1)
}

// Define extract
func (m *MaxHeap) Extract() int {
	extracted := m.array[0]
	lastIndex := len(m.array) - 1

	if len(m.array) == 0 {
		fmt.Printf("cant extract because heap is empty")
		return -1
	}

	m.array[0] = m.array[lastIndex]
	m.array = m.array[:lastIndex]

	m.maxHeapifyDown(0)

	return extracted
}

// Rearrange maxHeap up
func (m *MaxHeap) maxHeapifyUp(index int) {
	for m.array[parent(index)] < m.array[index] {
		m.swap(parent(index), index)
		index = parent(index)
	}
}

// Rearrange maxHeap down
func (m *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(m.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if m.array[l] > m.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if m.array[index] < m.array[childToCompare] {
			m.swap(index, childToCompare)
			index = childToCompare
			l, r = left(index), right(index)
		} else {
			return
		}
	}

}

// get parent index
func parent(index int) int {
	return (index - 1) / 2
}

// get the left child index
func left(parentIndex int) int {
	return 2*parentIndex + 1
}

// get the right child index
func right(parentIndex int) int {
	return 2*parentIndex + 2
}

// swap value for two index
func (m *MaxHeap) swap(i1 int, i2 int) {
	m.array[i1], m.array[i2] = m.array[i2], m.array[i1]
}

func main() {
	myHeap := &MaxHeap{}
	buildHeap := []int{12, 432, 26, 665, 50}
	for _, v := range buildHeap {
		myHeap.Insert(v)
		fmt.Printf("my heap is %v\n", myHeap)
	}
	myHeap.Extract()
	fmt.Printf("my extracted heap is %v\n", myHeap)
}
