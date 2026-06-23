// Безопасно выбросить» (discard invariant)
package main

import "fmt"

func twoSum(nums []int, target int) []int {
	left, right := 0, len(nums)-1

	for left < right {
		sum := nums[left] + nums[right]
		if sum == target {
			return []int{left + 1, right + 1}
		} else if sum < target {
			left++
		} else {
			right--
		}
	}
	return nil
}

func main() {
	nums := []int{2, 3, 4}
	fmt.Println(twoSum(nums, 6))
}
