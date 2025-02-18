package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(isAnagramNaive("racecar", "carrace"))
	fmt.Println(isAnagramEfficient("racecar", "carrace"))
}

// Time: O(n) Space: O(n)
func isAnagramEfficient(s, t string) bool {
	if len(s) != len(t) {
		return false
	}

	// Make frequency map
	charCount := make(map[rune]int)

	// Count all occurences of each character in s
	for _, char := range s {
		charCount[char]++
	}

	// Decrement occurences from s, using chars in t
	for _, char := range t {
		charCount[char]--

		// if any count goes negative, strings are not anagrams
		if charCount[char] < 0 {
			return false
		}
	}

	return true
}

// Time: O(n log n) Space: O(n)
func isAnagramNaive(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	sortedS, sortedT := sortStrings(s, t)
	return sortedS == sortedT
}

func sortStrings(s, t string) (string, string) {
	rS := []rune(s)
	rT := []rune(t)

	sort.Slice(rS, func(i, j int) bool {
		return rS[i] < rS[j]
	})

	sort.Slice(rT, func(i, j int) bool {
		return rT[i] < rT[j]
	})
	return string(rS), string(rT)
}
