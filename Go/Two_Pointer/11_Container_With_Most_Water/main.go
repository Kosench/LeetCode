package main

import "fmt"

func maxArea(height []int) int {
	var result int
	left, right := 0, len(height)-1

	for left < right {
		area := (right - left) * min(height[left], height[right])
		if area > result {
			result = area
		}
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	height := []int{1, 7, 2, 5, 4, 7, 3, 6}
	fmt.Println(maxArea(height))
}
