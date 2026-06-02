package main

import "fmt"

func sortedSquares(nums []int) []int {
	n := len(nums)
	res := make([]int, n)
	left, right := 0, n-1

	for i := n - 1; i >= 0; i-- {
		lSq := nums[left] * nums[left]
		rSq := nums[right] * nums[right]

		if lSq > rSq {
			res[i] = lSq
			left++
		} else {
			res[i] = rSq
			right--
		}
	}

	return res
}

func main() {
	fmt.Println(sortedSquares([]int{-4, -1, 0, 3, 10}))
	fmt.Println(sortedSquares([]int{-7, -3, 2, 3, 11}))
}
