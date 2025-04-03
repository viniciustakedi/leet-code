package main

import "fmt"

func containsDuplicate(nums []int) bool {
	m := make(map[int]int)
	containsDuplicate := false

	for _, num := range nums {
		_, ok := m[num]

		if !ok {
			m[num] = num
			continue
		}

		containsDuplicate = true
		break
	}

	return containsDuplicate
}

func main() {
	numArray := []int{0, 4, 5, 0, 3, 6}
	fmt.Print(containsDuplicate(numArray))
}
