package main

import (
	"fmt"
	"sort"
)

func topKFrequent(nums []int, k int) []int {
	counts := make(map[int]int, len(nums))
	for _, num := range nums {
		counts[num]++
	}

	type pair struct {
		num   int
		count int
	}
	var pairs []pair
	for num, count := range counts {
		pairs = append(pairs, pair{num, count})
	}

	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i].count > pairs[j].count
	})

	result := make([]int, 0, k)
	for i := 0; i < k && i < len(pairs); i++ {
		result = append(result, pairs[i].num)
	}

	return result
}

func main() {
	nums := []int{1, 1, 1, 2, 2, 3}
	k := 2
	fmt.Println(topKFrequent(nums, k))

}
