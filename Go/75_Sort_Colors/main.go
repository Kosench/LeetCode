package main

import "fmt"

func sortColors(nums []int) []int {
	low, mid := 0, 0
	high := len(nums) - 1

	for mid <= high {
		if nums[mid] == 0 {
			nums[low], nums[mid] = nums[mid], nums[low]
			low++
			mid++
		} else if nums[mid] == 2 {
			nums[high], nums[mid] = nums[mid], nums[high]
			high--
		} else {
			mid++
		}
	}
	return nums
}

func main() {
	fmt.Println(sortColors([]int{2, 0, 2, 1, 1, 0}))
}
