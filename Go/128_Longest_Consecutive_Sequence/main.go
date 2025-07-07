package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	numSet := make(map[int]bool)
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0
	for num := range numSet {
		if !numSet[num-1] {
			currentNum := num
			currentStreak := 1

			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
			}

			if currentStreak > longest {
				longest = currentStreak
			}
		}
	}

	return longest
}

func main() {
	intSl := []int{100, 4, 200, 1, 3, 2}
	intSl2 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	intSl3 := []int{1, 0, 1, 2}
	fmt.Println(longestConsecutive(intSl))
	fmt.Println(longestConsecutive(intSl2))
	fmt.Println(longestConsecutive(intSl3))
}
