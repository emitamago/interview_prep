package main

import "fmt"

// 1. Given an array of numbers sorted in ascending order and a target sum, find a pair in the array whose sum is equal to the given target.
type SolutionPairWithTargetSum struct{}

func (s SolutionPairWithTargetSum) pairWithTarget(arr []int, target int) []int {
	result := []int{-1, -1}
	first, last := 0, len(arr)-1
	for first < last {
		if arr[first]+arr[last] < target {
			first++
		} else if arr[first]+arr[last] > target {
			last--
		} else {
			result[0], result[1] = first, last
			break
		}
	}
	return result
}

// 2. Given an array of sorted numbers, move all non-duplicate number instances at the beginning of the array in-place. The relative order of the elements should be kept the same and you should not use any extra space so that the solution has constant space complexity i.e., .
type SolutionNexDuplicate struct{}

func (s *SolutionNexDuplicate) removeDuplicate(arr []int) int {
	//
	nextNonDuplicate := 1

	for i := 0; i < len(arr); i++ {
		if arr[nextNonDuplicate-1] != arr[i] {
			arr[nextNonDuplicate] = arr[i]
			nextNonDuplicate++
		}
	}
	return nextNonDuplicate
}

func main() {
	solution1 := &SolutionPairWithTargetSum{}
	pairNums1 := []int{1, 2, 3, 4, 6}
	target1 := 6
	pairNums2 := []int{2, 5, 9, 11}
	target2 := 11
	pairNums3 := []int{2, 5, 9}
	target3 := 3
	fmt.Printf("pair is %v\n", solution1.pairWithTarget(pairNums1, target1))
	fmt.Printf("pair is %v\n", solution1.pairWithTarget(pairNums2, target2))
	fmt.Printf("pair is %v\n", solution1.pairWithTarget(pairNums3, target3))

	solution2 := &SolutionNexDuplicate{}
	arr1 := []int{2, 3, 3, 3, 6, 9, 9}
	arr2 := []int{2, 2, 2, 11}
	fmt.Printf("non duplicate is %v\n", solution2.removeDuplicate(arr1))
	fmt.Printf("non duplicate is %v\n", solution2.removeDuplicate(arr2))

}
