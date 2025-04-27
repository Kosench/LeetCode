package __Median_of_Two_Sorted_Arrays

import "sort"

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2[:]...)
	sort.Ints(nums1)
	totalLength := len(nums1)

	if totalLength%2 == 0 {
		mid := totalLength / 2
		return float64(nums1[mid-1]+nums1[mid]) / 2.0
	}
	return float64(nums1[totalLength/2])

}
