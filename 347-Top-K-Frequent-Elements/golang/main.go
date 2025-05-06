package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	count := make(map[int]int)
	arrNums := []int{}

	for _, num := range nums {
		if _, ok := count[num]; !ok {
			arrNums = append(arrNums, num)
		}

		count[num]++
	}

	sort.Slice(arrNums, func(i, j int) bool {
		return count[arrNums[i]] > count[arrNums[j]]
	})

	var response []int

	i := 0
	for i < k {
		response = append(response, arrNums[i])
		i++
	}

	return response
}

func main() {
	arr := []int{1, 2, 2, 1, 3, 1, 2, 2, 3, 3, 3, 3, 2}
	fmt.Print(topKFrequent(arr, 2))
}
