package main

import "fmt"

// Create type queue
type Queue struct {
	items []int
}

// create enqueue
func (q *Queue) enqueue(i int) {
	q.items = append(q.items, i)
}

// create dequeue
func (q *Queue) dequeue() (i int) {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	myQueue.enqueue(3)
	myQueue.enqueue(6)
	myQueue.enqueue(70)
	fmt.Printf("my queue is %v\n", myQueue)
	myQueue.dequeue()
	fmt.Printf("my queue is %v\n", myQueue)
}
