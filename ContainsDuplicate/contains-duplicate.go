package main

import (
	"fmt"
)

type (
	Set      map[int]struct{}
	EmptyVal struct{}
)

func NewSet() Set {
	return Set{}
}

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 4, 1}))
}

func containsDuplicate(nums []int) bool {
	seen := NewSet()
	for _, num := range nums {
		if _, exists := seen[num]; exists {
			return true
		}
		seen[num] = EmptyVal{}
	}
	return false
}
