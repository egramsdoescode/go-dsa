package main

import "fmt"

func main() {
	fmt.Println(containsDuplicate([]int{1, 2, 3, 3}))
}

func containsDuplicate(nums []int) bool {
	// Create hashmap that acts as a set, use empty structs as the value to save memory
	seen := make(map[int]struct{})
	// Loop through the slice, check if the key exists, if it does return true
	for _, val := range nums {
		if _, exists := seen[val]; exists {
			return true
		}

		// Otherwise put the slice value as the key and empty struct as a value
		seen[val] = struct{}{}
	}
	return false
}
