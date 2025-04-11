package main

import (
	"fmt"
	"regexp"
	"strings"
)

func isPalindrome(s string) bool {
	nonAlphaNumericRegex := regexp.MustCompile(`[^a-zA-Z0-9]+`)
	s = strings.ToLower(nonAlphaNumericRegex.ReplaceAllString(s, ""))

	j := 0

	for i := len(s) - 1; i > 0; i-- {
		fmt.Println(s[j], s[i])
		if s[i] != s[j] {
			return false
		}
		j++
	}

	return true
}

func main() {
	// palindrome := "A man, a plan, a canal: Panama"
	palindrome := "race a car"

	fmt.Print(isPalindrome(palindrome))
}
