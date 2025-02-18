package main

import (
	"fmt"
	"sort"
)

func main() {
    fmt.Println(topKFrequentNaive([]int{1, 1, 1, 2, 2, 3}, 2))
    fmt.Println(topKFrequentEfficient([]int{1, 1, 1, 2, 2, 3}, 2))
}

// (Bucket Sort) Time: O(n) Space: O(n)
func topKFrequentEfficient(nums []int, k int) []int {
    freqMap := make(map[int]int)
    // Count frequency of each number 
    for _, num := range nums {
        freqMap[num]++
    }
    
    // Create buckets
    n := len(nums)
    buckets := make([][]int, n + 1)

    for num, freq := range freqMap {
        buckets[freq] = append(buckets[freq], num)
    }
    
    // Collect top k elements from highest frequency
    result := make([]int, 0, k)
    for i := n; i > 0; i-- {
        if len(buckets[i]) > 0 {
            result = append(result, buckets[i]...)
        }
    }

    return result[:k] 
}

// Time: O(n log n) Space: O(n)
func topKFrequentNaive(nums []int, k int) []int {
    freqMap := make(map[int]int)
    result := make([]int, 0, k)
    for _, value := range nums {
        freqMap[value]++
    }

    type kv struct {
        Key int
        Value int
    }

    var ss []kv
    for k, v := range freqMap {
        ss = append(ss, kv{k, v}) 
    }

    sort.Slice(ss, func(i, j int) bool {
        return ss[i].Value > ss[j].Value
    })
    
    for i := 0; i < k; i++ {
        result = append(result, ss[i].Key)
    }
    return result 
}
