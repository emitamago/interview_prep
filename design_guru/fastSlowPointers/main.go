package main

import (
	"fmt"
)

// 1. Given the head of a Singly LinkedList, write a function to determine if the LinkedList has a cycle in it or not.
type SolLinkedListCycle struct{}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (s SolLinkedListCycle) linkedListCycle(head *ListNode) bool {
	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next

		if slow == fast {
			return true
		}
	}
	return false
}

// 2. Given the head of a Singly LinkedList, write a method to return the middle node of the LinkedList.
type SolMiddleNode struct{}

func (s SolMiddleNode) findMiddleNode(head *ListNode) *ListNode {
	slow := head
	fast := head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

// 3. Given the head of a Singly LinkedList that contains a cycle, write a function to find the starting node of the cycle
type SolStartLinkedCycle struct{}

func (s SolStartLinkedCycle) findStart(head *ListNode, cycleLength int) *ListNode {
	pointer1 := head
	pointer2 := head

	// Move pointer2 for cycleLength
	for cycleLength > 0 {
		pointer2 = pointer2.Next
		cycleLength--
	}

	// move both pointers until they meet
	for pointer1 != pointer2 {
		pointer1 = pointer1.Next
		pointer2 = pointer2.Next
	}
	return pointer1

}

func (s SolStartLinkedCycle) findCycleCount(slow *ListNode) int {
	cycleLength := 0
	current := slow
	for {
		current = current.Next
		cycleLength++
		if current == slow {
			break
		}
	}
	return cycleLength
}

func (s SolStartLinkedCycle) findCycleStart(head *ListNode) *ListNode {
	cycleLenght := 0

	slow := head
	fast := head

	// find where slow and fast meet --- > that is after cycle start
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if fast == slow {
			cycleLenght = s.findCycleCount(slow)
			break
		}
	}
	return s.findStart(head, cycleLenght)
}

// 3. Any number will be called a happy number if, after repeatedly replacing it with a number equal to the sum of the square of all of its digits, leads us to the number 1. All other (not-happy) numbers will never reach 1. Instead, they will be stuck in a cycle of numbers that does not include 1.
type SolHappyNumber struct{}

func (s SolHappyNumber) findHappyNumber(num int) bool {
	slow := num
	fast := num
	for {
		slow = s.findSquareSum(slow)
		fast = s.findSquareSum(s.findSquareSum(fast))

		if slow == fast {
			break
		}
	}

	return slow == 1
}

func (s SolHappyNumber) findSquareSum(num int) int {
	sum := 0
	var digit int
	for num > 0 {
		digit = num % 10
		sum += digit * digit
		num /= 10
	}
	return sum
}

func main() {
	head1 := &ListNode{Val: 1}
	head1.Next = &ListNode{Val: 2}
	head1.Next.Next = &ListNode{Val: 3}
	head1.Next.Next.Next = &ListNode{Val: 4}
	head1.Next.Next.Next.Next = &ListNode{Val: 5}
	head1.Next.Next.Next.Next.Next = &ListNode{Val: 6}

	sol1 := &SolLinkedListCycle{}
	fmt.Printf("this is cycled linked list %v\n", sol1.linkedListCycle(head1))

	head1.Next.Next.Next.Next.Next.Next = head1.Next.Next
	fmt.Printf("this is cycled linked list %v\n", sol1.linkedListCycle(head1))

	head2 := &ListNode{Val: 1}
	head2.Next = &ListNode{Val: 2}
	head2.Next.Next = &ListNode{Val: 3}
	head2.Next.Next.Next = &ListNode{Val: 4}
	head2.Next.Next.Next.Next = &ListNode{Val: 5}
	head2.Next.Next.Next.Next.Next = &ListNode{Val: 6}
	sol2 := &SolMiddleNode{}
	fmt.Printf("Middle node is %v\n", sol2.findMiddleNode(head2).Val)
	head2.Next.Next.Next.Next.Next.Next = &ListNode{Val: 7}
	fmt.Printf("Middle node is %v\n", sol2.findMiddleNode(head2).Val)

	sol3 := &SolStartLinkedCycle{}
	head2.Next.Next.Next.Next.Next.Next = head2.Next.Next
	fmt.Printf("Link start here %v\n", sol3.findCycleStart(head2).Val)

	head2.Next.Next.Next.Next.Next.Next = head2
	fmt.Printf("Link start here %v\n", sol3.findCycleStart(head2).Val)

	sol4 := &SolHappyNumber{}
	fmt.Printf("Is it happy number %v\n", sol4.findHappyNumber(23))
	fmt.Printf("Is it happy number %v\n", sol4.findHappyNumber(12))
}
