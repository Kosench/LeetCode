package main

func twoSum_Map(nums []int, target int) []int {
	valueMap := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		value := target - nums[i]
		if value, ok := valueMap[value]; ok {
			return []int{value, i}
		}
		valueMap[nums[i]] = i
	}
	return nil
}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return nil
}

func main() {
	nums := []int{2, 7, 11, 15}
	twoSum(nums, 9)
}
