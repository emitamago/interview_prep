package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) deleteWithValue(value int) {
	if l.length == 0 {
		return
	}
	if l.head.data == value {
		l.head = l.head.next
		l.length--
	}
	previousToDelete := l.head
	for previousToDelete.next.data != value {
		if previousToDelete.next.next == nil {
			return
		}
		previousToDelete = previousToDelete.next
	}
	previousToDelete.next = previousToDelete.next.next
	l.length--
}

func (l linkedList) printData() {
	toPrint := l.head
	for l.length != 0 {
		fmt.Printf("%d\n", toPrint.data)
		toPrint = toPrint.next
		l.length--
	}
	fmt.Printf("\n")
}

func main() {
	mylist := linkedList{}
	node1 := &node{data: 64}
	node2 := &node{data: 90}
	node3 := &node{data: 98}
	node4 := &node{data: 115}
	node5 := &node{data: 356}
	node6 := &node{data: 982}
	mylist.prepend(node1)
	mylist.prepend(node2)
	mylist.prepend(node3)
	mylist.prepend(node4)
	mylist.prepend(node5)
	mylist.prepend(node6)
	mylist.printData()
	mylist.deleteWithValue(115)
	mylist.printData()
	mylist.deleteWithValue(64)

	fmt.Printf("my list is %v\n", mylist)
}
