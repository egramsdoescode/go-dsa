package main

import (
	"fmt"
	"sort"
)

func main() {
	strs := []string{"eat", "tea", "tan", "ate", "nat", "bat"}
	fmt.Println(groupAnagramsNaive(strs))
}

func groupAnagramsEfficient(strs []string) [][]string {
	// TODO: Implement
	return nil
}

func groupAnagramsNaive(strs []string) [][]string {
	result := make([][]string, 0)
	anagrams := make(map[string][]string)

	for _, str := range strs {
		sortedStr := sortString(str)
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}

	for _, value := range anagrams {
		result = append(result, value)
	}

	return result
}

func sortString(str string) string {
	rStr := []rune(str)
	sort.Slice(rStr, func(i, j int) bool {
		return rStr[i] < rStr[j]
	})
	return string(rStr)
}
