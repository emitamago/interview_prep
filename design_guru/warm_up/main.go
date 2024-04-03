package main

import (
	"fmt"
	"math"
	"sort"
	"strings"
	"unicode"
)

// 1. Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.
type SolutionDuplicate struct{}

// Create hashmap
func (s SolutionDuplicate) countDuplicateHash(nums []int) bool {
	// create map for each int in array
	set := make(map[int]int)
	for _, x := range nums {
		if _, ok := set[x]; ok {
			return true
		}
		set[x] = 1
	}
	return false
}

// Use sort. only need to check next int
func (s SolutionDuplicate) countDuplicateSort(nums []int) bool {
	// sort nums
	sort.Ints(nums)
	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == nums[i+1] {
			return true
		}
	}
	return false
}

// 2. Given a string sentence containing English letters (lower or upper-case), return true if sentence is a pangram, or false otherwise.
type SolutionPangram struct{}

func (s SolutionPangram) checkPangram(sentence string) bool {
	set := make(map[rune]bool)
	for _, v := range strings.ToLower(sentence) {
		if unicode.IsLetter(v) {
			set[v] = true
		}
	}
	return len(set) == 26
}

// 3. Given a string s, reverse only all the vowels in the string and return it.
type SolutionReverseVowels struct{}

const vowels = "aeiouAEIOU"

// Two pointer solution --- move pointer until finding char we want
func (s SolutionReverseVowels) reverseVowels(word string) string {
	first, last := 0, len(word)-1
	array := []rune(word)
	for first < last {
		// loop until it finds first vowel
		for first < last && !strings.ContainsRune(vowels, array[first]) {
			first++
		}
		// loop until it finds last vowel
		for first < last && !strings.ContainsRune(vowels, array[last]) {
			last--
		}
		array[first], array[last] = array[last], array[first]
		first++
		last--
	}
	return string(array)
}

// 4. Given a string s, return true if it is a palindrome, or false otherwise
type SolutionPalindrome struct{}

// Two pointer solution ---- move pointer until finding char we want
func (s SolutionPalindrome) checkPalindrome(sentence string) bool {
	i, j := 0, len(sentence)-1
	for i < j {
		for i < j && !isLetterOrDigit(rune(sentence[i])) {
			i++
		}
		for i < j && !isLetterOrDigit(rune(sentence[j])) {
			j--
		}
		if strings.ToLower(string(sentence[i])) != strings.ToLower(string(sentence[j])) {
			return false
		}
		i++
		j--
	}
	return true
}

// Support function
func isLetterOrDigit(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r)
}

// 5. Given two strings s and t, return true if t is an anagram of s, and false otherwise.
type SolutionIsAnagram struct{}

// create one map for a word and increment by one per each occurance. and go through b and decrement each occurance
func (s SolutionIsAnagram) isAnagram(a, b string) bool {
	if len(a) != len(b) {
		return false
	}

	freMap := make(map[rune]int)
	for _, v1 := range a {
		freMap[v1]++
	}
	for _, v2 := range b {
		freMap[v2]--
	}

	for _, v3 := range freMap {
		if v3 != 0 {
			return false
		}
	}
	return true

}

// 6. Given an array of strings words and two different strings that already exist in the array word1 and word2, return the shortest distance between these two words in the list.
type SolutionShortestDistance struct{}

func (s SolutionShortestDistance) checkShortestDistance(words []string, w1 string, w2 string) int {
	// create position for w1 and w2 (-1 means the word has not been occure yet)
	p1, p2 := -1, -1
	shortestDistance := len(words)

	// iterate over words list to find position. if it find word1 upate p1 if word2 then update p2
	for i, w := range words {
		if w == w1 {
			p1 = i
		} else if w == w2 {
			p2 = i
		}
	}

	// if p1 and p2 are not negative then caluculate position and find if original distance or updated position are shorter
	if p1 != -1 && p2 != -1 {
		shortestDistance = int(math.Min(float64(shortestDistance), float64(math.Abs(float64(p1-p2)))))
	}

	return shortestDistance
}

// 7. Given an array of integers nums, return the number of good pairs
type SolutionGoodPair struct{}

func (s SolutionGoodPair) numOfGoodPair(nums []int) int {
	pair := 0
	numMap := make(map[int]int)
	for _, v := range nums {
		// increament each occurance of the number
		numMap[v]++

		// if a number occur 2 times. we can make pair one time so add numMap[v] - 1  to total pair
		pair += numMap[v] - 1
	}
	return pair
}

// 8. Given a non-negative integer x, return the square root of x rounded down to the nearest integer. The returned integer should be non-negative as well.
type SolutionSquare struct{}

func (s SolutionSquare) findSquare(x int) int {
	if x < 2 {
		return x
	}
	// 8
	left, right := 2, x/2 // 2 , 4
	var num int64
	var pivot int
	for left <= right { // binary search
		pivot = left + (right-left)/2     // pivot 2 + (4-2)/2 = 3
		num = int64(pivot) * int64(pivot) // 6
		if num > int64(x) {
			right = pivot - 1
		} else if num < int64(x) {
			left = pivot + 1
		} else {
			return pivot
		}
	}
	return right
}

func main() {
	solution1 := &SolutionDuplicate{}
	nums := []int{1, 2, 3, 4}
	nums2 := []int{1, 3, 2, 1}
	fmt.Printf("solution1 is %v\n", solution1.countDuplicateHash(nums))
	fmt.Printf("solution2 is %v\n", solution1.countDuplicateHash(nums2))
	fmt.Printf("solution1 is %v\n", solution1.countDuplicateSort(nums))
	fmt.Printf("solution2 is %v\n", solution1.countDuplicateSort(nums2))

	solution2 := &SolutionPangram{}
	sentence := "TheQuickBrownFoxJumpsOverTheLazyDog"
	sentence2 := "tttttttttttttttttttttttttt"
	sentence3 := "amg"
	fmt.Printf("solution1 is %v\n", solution2.checkPangram(sentence))
	fmt.Printf("solution2 is %v\n", solution2.checkPangram(sentence2))
	fmt.Printf("solution1 is %v\n", solution2.checkPangram(sentence3))

	solution3 := &SolutionReverseVowels{}
	s1 := "hello"
	s2 := "DesignGUrus"
	s3 := "bonbon"
	fmt.Printf("s1 is %v\n", solution3.reverseVowels(s1))
	fmt.Printf("s2 is %v\n", solution3.reverseVowels(s2))
	fmt.Printf("s3 is %v\n", solution3.reverseVowels(s3))

	solution4 := &SolutionPalindrome{}
	p1 := "A man, a plan, a canal, Panama!"
	p2 := "race a car"
	p3 := "q"
	fmt.Printf("p1 is %v\n", solution4.checkPalindrome(p1))
	fmt.Printf("p2 is %v\n", solution4.checkPalindrome(p2))
	fmt.Printf("p3 is %v\n", solution4.checkPalindrome(p3))

	solution5 := &SolutionIsAnagram{}
	a1, a2 := "listen", "silent"
	a3, a4 := "hello", "world"
	a5, a6 := "aha", "aeee"
	fmt.Printf("this is anagram: %v\n", solution5.isAnagram(a1, a2))
	fmt.Printf("this is anagram: %v\n", solution5.isAnagram(a3, a4))
	fmt.Printf("this is anagram: %v\n", solution5.isAnagram(a5, a6))

	solution6 := &SolutionShortestDistance{}
	words1 := []string{"the", "quick", "brown", "fox", "jumps", "over", "the", "lazy", "dog"}
	word11 := "fox"
	word21 := "dog"
	words3 := []string{"a", "c", "d", "b", "a"}
	word13 := "a"
	word23 := "b"
	fmt.Printf("distance is %v\n", solution6.checkShortestDistance(words1, word11, word21))
	fmt.Printf("distance is %v\n", solution6.checkShortestDistance(words3, word13, word23))

	solution7 := &SolutionGoodPair{}
	pair1 := []int{1, 2, 3, 1, 1, 3}
	pair2 := []int{1, 1, 1, 1}
	pair3 := []int{1, 2, 3}
	fmt.Printf("pair is %v\n", solution7.numOfGoodPair(pair1))
	fmt.Printf("pair is %v\n", solution7.numOfGoodPair(pair2))
	fmt.Printf("pair is %v\n", solution7.numOfGoodPair(pair3))

	solution8 := &SolutionSquare{}
	input1 := 4
	input2 := 8
	fmt.Printf("sqrt is %v\n", solution8.findSquare(input1))
	fmt.Printf("sqrt is %v\n", solution8.findSquare(input2))
}
