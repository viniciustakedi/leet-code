package main

import "fmt"

func twoSum(nums []int, target int) []int {
	var result []int
	for i, vOne := range nums {
		for j, vTwo := range nums {
			if i == j {
				continue
			}

			if (vOne + vTwo) == target {
				result = []int{i, j}
				break
			}

			if len(result) > 0 {
				break
			}
		}
	}

	return result[:]
}

func main() {
	arr := []int{3, 2, 4}

	fmt.Print(twoSum(arr, 6))
}
