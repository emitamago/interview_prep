package main

import (
	"fmt"
	"sort"
)

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

// 3. Given a sorted array, create a new array containing squares of all the numbers of the input array in the sorted order.
type SolutionSortedSquaredArray struct{}

func (s SolutionSortedSquaredArray) squareAndSort(snums []int) []int {
	n := len(snums)
	squaredSorted := make([]int, n)
	left, right := 0, n-1
	highestNumIndex := n - 1
	for left <= right {
		leftSquare := snums[left] * snums[left]
		rightSquare := snums[right] * snums[right]
		if leftSquare > rightSquare {
			squaredSorted[highestNumIndex] = leftSquare
			highestNumIndex--
			left++
		} else {
			squaredSorted[highestNumIndex] = rightSquare
			highestNumIndex--
			right--
		}

	}
	return squaredSorted
}

// 4. Given an array of unsorted numbers, find all unique triplets in it that add up to zero.
type SolutionTriSumZero struct{}

func (s *SolutionTriSumZero) tripletToZero(unums []int) [][]int {
	triplet := make([][]int, 0)
	sort.Ints(unums)
	for i := 0; i < len(unums)-2; i++ {
		if i > 0 && unums[i] == unums[i-1] {
			continue
		}
		s.searchPair(unums, -unums[i], i+1, &triplet) // left index is next element of current element unums[i]

	}
	return triplet
}

func (s *SolutionTriSumZero) searchPair(arr []int, targetSum int, left int, triplet *[][]int) {
	right := len(arr) - 1 // last element of original array
	for left < right {
		currentSum := arr[left] + arr[right]
		if currentSum == targetSum {
			*triplet = append(*triplet, []int{-targetSum, arr[right], arr[left]})
			left++
			right--
			for left < right && arr[left] == arr[left-1] {
				left++
			}
			for left < right && arr[right] == arr[right+1] {
				right--
			}
		} else if currentSum < targetSum {
			left++
		} else {
			right--
		}
	}
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

	solution3 := &SolutionSortedSquaredArray{}
	snums1 := []int{-2, -1, 0, 2, 3}
	snums2 := []int{-3, -1, 0, 1, 2}
	fmt.Printf("square and sorted array is %v\n", solution3.squareAndSort(snums1))
	fmt.Printf("square and sorted array is %v\n", solution3.squareAndSort(snums2))

	solution4 := &SolutionTriSumZero{}
	unums1 := []int{-3, 0, 1, 2, -1, 1, -2}
	unums2 := []int{-5, 2, -1, -2, 3}
	fmt.Printf("triplets are %v\n", solution4.tripletToZero(unums1))
	fmt.Printf("triplets are %v\n", solution4.tripletToZero(unums2))

}
