package main

import (
	"fmt"
)

// Dynamic programming
// func maxProfit(prices []int) int {
// 	maxP := 0
// 	minBuy := math.MaxInt32

// 	for _, price := range prices {
// 		if price-minBuy > maxP {
// 			maxP = price - minBuy
// 		}

// 		if price < minBuy {
// 			minBuy = price
// 		}
// 	}

// 	return maxP
// }

// Two pointers
func maxProfit(prices []int) int {
	maxP := 0
	l, r := 0, 1

	for r < len(prices) {
		if prices[l] < prices[r] {
            minus := prices[r]-prices[l]
			if minus > maxP {
				maxP = minus
			}
		} else {
			l = r
		}
		r++
	}

	return maxP
}

func main() {
	arr := []int{7, 1, 5, 3, 6, 4} //5
	fmt.Println(maxProfit(arr))

	arr1 := []int{7, 6, 4, 3, 1} //0
	fmt.Println(maxProfit(arr1))

	arr2 := []int{2, 1, 4} //3
	fmt.Println(maxProfit(arr2))
}
