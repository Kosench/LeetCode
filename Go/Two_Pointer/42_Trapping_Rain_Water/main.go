// Два указателя с внешним максимумом-ограничителем
package main

import (
	"fmt"
)

func trap(height []int) int {
	maxL := make([]int, len(height))
	maxR := make([]int, len(height))

	current_max := 0
	for i := 0; i < len(height); i++ {
		current_max = max(current_max, height[i])
		maxL[i] = current_max
	}

	current_max = 0
	for i := len(height) - 1; i >= 0; i-- {
		current_max = max(current_max, height[i])
		maxR[i] = current_max
	}

	total := 0
	for i := 0; i < len(height); i++ {
		total += min(maxL[i], maxR[i]) - height[i]
	}

	return total
}

func trap_2(height []int) int {
	left, right := 0, len(height)-1
	maxL, maxR := 0, 0
	total := 0

	for left < right {
		maxL = max(maxL, height[left])
		maxR = max(maxR, height[right])

		if maxL < maxR {
			total += maxL - height[left]
			left++
		} else if maxR <= maxL {
			total += maxR - height[right]
			right--
		}

	}

	return total
}

func main() {
	fmt.Println(trap_2([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))
}
