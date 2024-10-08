package search

import "fmt"

func BinarySearch() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := Binary(nums, 5)
	if !result {
		fmt.Printf("target not found in slice")
	} else {
		fmt.Printf("target found in slice")
	}
}

func Binary(nums []int, target int) bool {
	start := 0
	end := len(nums) - 1

	for start <= end {
		mid := (start + end) / 2
		valAtMid := nums[mid]

		if valAtMid == target {
			return true
		} else if valAtMid > target {
			end = mid
		} else if valAtMid < target {
			start = mid + 1
		}
	}

	return false
}
