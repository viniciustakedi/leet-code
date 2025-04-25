package main

import "fmt"

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		area := min(height[left], height[right]) * (right - left)
		if area > maxArea {
			maxArea = area
		}

		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
}

func main() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(height))
	height = []int{1, 1}
	fmt.Println(maxArea(height))
}
