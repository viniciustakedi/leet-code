package main

import (
	"fmt"
)

func isValid(s string) bool {
	stack := []rune{}
	closeToOpen := map[rune]rune{')': '(', ']': '[', '}': '{'}

	for _, c := range s {
		if open, exists := closeToOpen[c]; exists {
			if len(stack) > 0 {
				if stack[len(stack)-1] != open {
					return false
				}
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		} else {
			stack = append(stack, c)
		}
	}

	return len(stack) == 0
}

func main() {
	str := "[([]])"

	fmt.Print(isValid(str))
}
