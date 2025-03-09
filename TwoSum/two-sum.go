package main

import "fmt"

func main() {
	fmt.Println(twoSum([]int{1, 2, 3, 4}, 7))
}

// Time: O(n) Space: O(n)
func twoSum(nums []int, target int) []int {
	seen := make(map[int]int)
	for index, value := range nums {
		complement := target - value
		if idx, exists := seen[complement]; exists {
			return []int{idx, index}
		}
		seen[value] = index
	}

	return nil
}
