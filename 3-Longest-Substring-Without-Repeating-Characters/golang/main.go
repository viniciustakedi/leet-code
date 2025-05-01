package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	hash := make(map[byte]bool)
	l, res := 0, 0

	for r := range s {
		for hash[s[r]] {
			delete(hash, s[l])
			l++
		}

		hash[s[r]] = true
		if r-l+1 > res {
			res = r - l + 1
		}
	}

	return res
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"), 3)
	fmt.Println(lengthOfLongestSubstring("bbbbb"), 1)
	fmt.Println(lengthOfLongestSubstring("pwwkew"), 3)
	fmt.Println(lengthOfLongestSubstring("dvdf"), 3)
}
