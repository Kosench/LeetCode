package main

import (
	"fmt"
	"sort"
)

// map[int]struct{} — это позволит избежать лишнего хранения bool значений
func containsDuplicate_2(nums []int) bool {
	m := make(map[int]struct{})
	for _, val := range nums {
		if _, exists := m[val]; exists {
			return true
		}
		m[val] = struct{}{}
	}
	return false
}

func containsDuplicate(nums []int) bool {
	m := make(map[int]bool)
	for _, val := range nums {
		if m[val] {
			return true
		}
		m[val] = true
	}
	return false
}

func containsDuplicate_sort(nums []int) bool {
	sort.Ints(nums) // O(n log n)
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			return true
		}
	}
	return false
}

func main() {
	arr := []int{1, 2, 3, 4}
	fmt.Println(containsDuplicate(arr))
}
