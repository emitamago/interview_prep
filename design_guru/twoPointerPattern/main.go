package main

import (
	"fmt"
	"math"
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
		searchPair(unums, -unums[i], i+1, &triplet) // left index is next element of current element unums[i]

	}
	return triplet
}

func searchPair(arr []int, targetSum int, left int, triplet *[][]int) {
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

// 5. Given an array of unsorted numbers and a target number, find a triplet in the array whose sum is as close to the target number as possible, return the sum of the triplet. If there are more than one such triplet, return the sum of the triplet with the smallest sum
type SolutionTripletSumClose struct{}

func (s SolutionTripletSumClose) tripletCloseToSum(arr []int, targetSum int) int {
	if arr == nil || len(arr) < 3 {
		panic("IllegalArgumentException")
	}

	sort.Ints(arr)
	smallestDifference := math.MaxInt32
	for i := 0; i < len(arr)-2; i++ {
		left, right := i+1, len(arr)-1
		for left < right {
			// finding the target difference
			targetDiff := targetSum - arr[i] - arr[left] - arr[right]
			if targetDiff == 0 { // found a triplet with an exact sum
				return targetSum // return sum of all the numbers
			}

			// handle the smallest sum when we have more than one solution
			if math.Abs(float64(targetDiff)) < math.Abs(float64(smallestDifference)) ||
				(math.Abs(float64(targetDiff)) == math.Abs(float64(smallestDifference)) &&
					targetDiff > smallestDifference) {
				smallestDifference = targetDiff // save the closest and the biggest difference
			}

			if targetDiff > 0 {
				left++ // need a triplet with a bigger sum
			} else {
				right-- // need a triplet with a smaller sum
			}
		}
	}
	return targetSum - smallestDifference
}

// 6. Given an array arr of unsorted numbers and a target sum, count all triplets in it such that arr[i] + arr[j] + arr[k] < target where i, j, and k are three different indices. Write a function to return the count of such triplets.
type SolutionTripletSmallerSum struct{}

func (s SolutionTripletSmallerSum) tripleWithSmallerSum(arr []int, target int) int {
	if arr == nil || len(arr) < 3 {
		return 0
	}
	count := 0
	sort.Ints(arr)
	for i := 0; i < len(arr)-2; i++ {
		count += s.searchSmallerPair(arr, target-arr[i], i)
	}
	return count
}

func (s SolutionTripletSmallerSum) searchSmallerPair(arr []int, targetSum int, first int) int {
	count := 0
	left, right := first+1, len(arr)-1
	for left < right {
		if arr[left]+arr[right] < targetSum { // found the triplet
			// since arr[right] >= arr[left], therefore, we can replace arr[right] by any
			// number between left and right to get a sum less than the target sum
			count += right - left
			left++
		} else {
			right-- // we need a pair with a smaller sum
		}
	}
	return count
}

// 7. Given an array with positive numbers and a positive target number, find all of its contiguous subarrays whose product is less than the target number.
type SolutionSubarrayProductLessThanTaget struct{}

func (a SolutionSubarrayProductLessThanTaget) subarraysProduct(arr []int, target int) [][]int {
	var result [][]int
	product := 1
	left := 0

	for right := 0; right < len(arr); right++ {
		product *= arr[right]

		for product >= target && left < len(arr) {
			product /= arr[left]
			left++
		}

		var tempList []int

		for i := right; i >= left; i-- {
			tempList = append([]int{arr[i]}, tempList...)

			result = append(result, append([]int(nil), tempList...))
		}
	}
	return result
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

	solution5 := &SolutionTripletSumClose{}
	closeNums1 := []int{-1, 0, 2, 3}
	closeNum1 := 2
	closeNums2 := []int{-3, -1, 1, 2}
	closeNum2 := 1
	fmt.Printf("closest diff is %v\n", solution5.tripletCloseToSum(closeNums1, closeNum1))
	fmt.Printf("closest diff is %v\n", solution5.tripletCloseToSum(closeNums2, closeNum2))

	solution6 := &SolutionTripletSmallerSum{}
	smallnums1 := []int{-1, 0, 2, 3}
	smallnum1 := 3
	smallnums2 := []int{-1, 4, 2, 1, 3}
	smallnum2 := 5
	fmt.Printf("smallest triplet is %v\n", solution6.tripleWithSmallerSum(smallnums1, smallnum1))
	fmt.Printf("smallest triplet is %v\n", solution6.tripleWithSmallerSum(smallnums2, smallnum2))

	solution7 := &SolutionSubarrayProductLessThanTaget{}
	productNums1 := []int{2, 5, 3, 10}
	productTarget1 := 30
	productNums2 := []int{8, 2, 6, 5}
	productTarget2 := 50
	fmt.Printf("product subarray is %v\n", solution7.subarraysProduct(productNums1, productTarget1))
	fmt.Printf("product subarray is %v\n", solution7.subarraysProduct(productNums2, productTarget2))

}
