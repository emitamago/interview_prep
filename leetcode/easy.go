package main

import "fmt"

// Input: nums = [1,2,3,4,5,6,7], k = 3
// Output: [5,6,7,1,2,3,4]
// Explanation:
// rotate 1 steps to the right: [7,1,2,3,4,5,6]
// rotate 2 steps to the right: [6,7,1,2,3,4,5]
// rotate 3 steps to the right: [5,6,7,1,2,3,4]

func rotate(nums []int, k int) []int {
	step := 0
	for step <= k {
		copy(nums, append(nums[len(nums)-k:], nums[:len(nums)-k]...))
		step++
	}
	return nums
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	steps := 3
	newNums := rotate(nums, steps)
	fmt.Printf("new list is :%v\n", newNums)
}
