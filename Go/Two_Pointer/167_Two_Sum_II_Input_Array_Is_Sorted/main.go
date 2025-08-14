package main

import "fmt"

func twoSum_2(numbers []int, target int) []int { //Hash Map
	// map: число → индекс (0-based)
	numMap := make(map[int]int)

	for i, num := range numbers {
		complement := target - num
		if j, ok := numMap[complement]; ok {
			// нашли пару
			return []int{j + 1, i + 1} // 1-based индексы
		}
		numMap[num] = i
	}

	return nil
}

func twoSum(numbers []int, target int) []int { // Two Pointer
	left, right := 0, len(numbers)-1
	for left < right {
		sum := numbers[left] + numbers[right]
		if sum == target {
			return []int{left + 1, right + 1} // 1-based
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
