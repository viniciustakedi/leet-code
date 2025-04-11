package main

import (
	"fmt"
)

// First attempt was a susscess but the complexity stayed in O(n log n), and
// the leet code problem requires a O(n) solution.
// func longestConsecutive(nums []int) int {
// 	sort.Slice(nums, func(i, j int) bool {
// 		return nums[i] < nums[j]
// 	})

// 	var lastNum int
// 	biggerSequence := 1
// 	currentSequence := 1

// 	arrLength := len(nums)

// 	if arrLength <= 0 {
// 		biggerSequence = 0
// 	} else {
// 		lastNum = nums[0]
// 	}

// 	for i, num := range nums {
// 		if i == 0 || num == lastNum {
// 			continue
// 		}

// 		if lastNum+1 == num {
// 			currentSequence++
// 		} else {
// 			currentSequence = 1
// 		}

// 		if currentSequence > biggerSequence {
// 			biggerSequence = currentSequence
// 		}

// 		lastNum = num
// 	}

// 	return biggerSequence
// }

// This attempt call the requirements of leet code problem, even result is
// slower the complexity is O(n)
func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longestSequence := 0

	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longestSequence {
				longestSequence = currentStreak
			}
		}
	}

	return longestSequence
}

func main() {
	// arr := []int{100, 4, 200, 1, 3, 2}
	arr := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	// arr := []int{1, 0, 1, 2}

	fmt.Print(longestConsecutive(arr))
}
