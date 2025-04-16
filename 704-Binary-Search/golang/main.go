package main

import "fmt"

// This function works very well! O got a O(log n) that problem requires, but I think I can make better.
// func search(nums []int, target int) int {
// 	currentIndex := int(len(nums) / 2)
// 	currentArr := nums

// 	mapIdx := make(map[int]int)
// 	for i, v := range nums {
// 		mapIdx[v] = i
// 	}

// 	end := false
// 	for !end {
// 		if len(currentArr) == 1 {
// 			end = true

// 			if currentArr[0] == target {
// 				return mapIdx[currentArr[0]]
// 			} else {
// 				return -1
// 			}
// 		}

// 		value := currentArr[currentIndex]
// 		if currentArr[currentIndex] == target {
// 			end = true
// 			return mapIdx[currentArr[currentIndex]]
// 		}

// 		if value < target {
// 			currentArr = currentArr[currentIndex:]
// 		} else {
// 			currentArr = currentArr[:currentIndex]
// 		}

// 		currentIndex = int(len(currentArr) / 2)
// 	}

// 	return -1
// }

func search(nums []int, target int) int {
	l, r := 0, len(nums)-1

	for l <= r {
		m := int((l + r) / 2)
		if nums[m] > target {
			r = m - 1
		} else if nums[m] < target {
			l = m + 1
		} else {
			return m
		}
	}

	return -1
}

func main() {
	arr := []int{-1, 0, 3, 5, 9, 12}
	fmt.Print(search(arr, 9))
}
