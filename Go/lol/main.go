package main

import (
	"fmt"
)

func max_podryd(nums []int) int {
	max_len := 1
	temp := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == nums[i-1] {
			temp++
		} else {
			if temp > max_len {
				max_len = temp
			}
			temp = 1
		}
	}

	if temp > max_len {
		max_len = temp
	}

	return max_len
}

func max_sled(nums []int) {
	mapNum := make(map[int]int, len(nums))

	for _, num := range nums {
		if _, exist := mapNum[num]; !exist {
			mapNum[num] = 1
		} else {
			mapNum[num]++
		}
	}
	maxNum := 0
	for _, j := range mapNum {
		if j > maxNum {
			maxNum = j
		}
	}
	fmt.Println(maxNum)
}

func main() {
	//max_sled([]int{1, 1, 1, 4, 6, 7, 8, 4, 4, 4, 4})
	fmt.Println(max_podryd([]int{1, 1, 1, 4, 6, 7, 8, 4, 4, 4, 4}))
}
