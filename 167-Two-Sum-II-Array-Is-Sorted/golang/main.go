package main

import "fmt"

func twoSum(numbers []int, target int) []int {
	numsMap := make(map[int][]int)
	var response []int

	for i, n := range numbers {
		numsMap[n] = append(numsMap[n], i)
	}

	for i, n := range numbers {
		var key int
		if target > n || n < 0 {
			key = target - n
		} else {
			key = n - target
		}

		if v, ok := numsMap[key]; ok {
			if len(v) <= 0 {
				continue
			}

			sndIndex := v[i]
			if len(v) > 1 {
				sndIndex++
			}

			response = append(response, i+1, sndIndex+1)
			break
		}
	}

	return response
}

func main() {
	// arr := []int{2, 7, 11, 15} // t=9
	// arr := []int{2, 3, 4} // t=6
	// arr := []int{-1, 0} // t=-1
	arr := []int{-1, -1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}

	fmt.Print(twoSum(arr, -2))
}
