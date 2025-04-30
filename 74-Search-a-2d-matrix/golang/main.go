package main

import (
	"fmt"
)

func searchMatrix(matrix [][]int, target int) bool {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return false
	}

	top, bottom := 0, len(matrix)-1
	for top <= bottom {
		mid := (top + bottom) / 2
		if target >= matrix[mid][0] && target <= matrix[mid][len(matrix[mid])-1] {
			return binarySearch(matrix[mid], target)
		} else if target < matrix[mid][0] {
			bottom = mid - 1
		} else {
			top = mid + 1
		}
	}

	return false
}

func binarySearch(row []int, target int) bool {
	left, right := 0, len(row)-1

	for left <= right {
		mid := (left + right) / 2
		if row[mid] == target {
			return true
		} else if row[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return false
}

func main() {
	// matrix := [][]int{
	// 	{1, 3, 5, 7},
	// 	{10, 11, 16, 20},
	// 	{23, 30, 34, 60},
	// }

	// fmt.Println(searchMatrix(matrix, 3))
	// fmt.Println(searchMatrix(matrix, 13))

	matrix := [][]int{
		{1},
		{3},
	}

	fmt.Println(searchMatrix(matrix, 2))

	// matrix = [][]int{{1, 3, 5}}
	// fmt.Println(searchMatrix(matrix, 3))
}
