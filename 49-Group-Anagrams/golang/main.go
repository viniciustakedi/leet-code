package main

import (
	"fmt"
	"sort"
)

// This first attempt I achieved the result expected, but with huge arrays the code was very slow.
// func groupAnagrams(strs []string) [][]string {
// 	arrContainsOneElement := len(strs) == 1

// 	if len(strs) <= 0 || (arrContainsOneElement && strs[0] == "") {
// 		return [][]string{{""}}
// 	}

// 	if arrContainsOneElement {
// 		return [][]string{{strs[0]}}
// 	}

// 	var response [][]string
// 	var currentMap map[rune]int
// 	var currentAnagrams []string

// 	strAlreadyChecked := []string{}

// 	for i, v := range strs {
// 		if slices.Contains(strAlreadyChecked, v) {
// 			continue
// 		}

// 		var isAnAnagram bool
// 		currentAnagrams = []string{v}

// 		for iv, cv := range strs {
// 			currentMap = make(map[rune]int)
// 			isAnAnagram = true

// 			if v != "" && cv == "" {
// 				continue
// 			}

// 			if i == iv {
// 				continue
// 			}

// 			if len(v) == len(cv) {
// 				for _, char := range v {
// 					currentMap[char]++
// 				}

// 				for _, char := range cv {
// 					currentMap[char]--

// 					if currentMap[char] < 0 {
// 						isAnAnagram = false
// 					}
// 				}
// 			} else {
// 				isAnAnagram = false
// 			}

// 			if isAnAnagram {
// 				currentAnagrams = append(currentAnagrams, cv)

// 				if !slices.Contains(strAlreadyChecked, cv) {
// 					strAlreadyChecked = append(strAlreadyChecked, cv)
// 				}

// 				continue
// 			}
// 		}

// 		if !slices.Contains(strAlreadyChecked, v) {
// 			strAlreadyChecked = append(strAlreadyChecked, v)
// 		}

// 		if len(currentAnagrams) == 1 {
// 			response = append([][]string{{v}}, response...)
// 		} else {
// 			slices.SortFunc(currentAnagrams, func(a, b string) int {
// 				return strings.Compare(strings.ToLower(a), strings.ToLower(b))
// 			})

// 			response = append(response, currentAnagrams)
// 		}
// 	}

// 	slices.SortFunc(response, func(a, b []string) int {
// 		return len(a) - len(b)
// 	})

// 	return response
// }

func groupAnagrams(strs []string) [][]string {
	anagramMap := make(map[string][]string)

	var key string
	for _, str := range strs {
		runes := []rune(str)

		sort.Slice(runes, func(i, j int) bool {
			return runes[i] < runes[j]
		})

		key = string(runes)
		anagramMap[key] = append(anagramMap[key], str)
	}

	result := make([][]string, 0, len(anagramMap))
	for _, group := range anagramMap {
		result = append(result, group)
	}

	return result
}

func main() {
	// arr := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	arr := []string{"", ""}
	fmt.Print(groupAnagrams(arr))
}
