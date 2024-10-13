// Chapter 2: Why Algorithms Matter
package search

import "fmt"

func BinarySearch() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Binary Search:")
	result := Binary(nums, 5)
	if !result {
		fmt.Printf("target not found in slice\n\n")
	} else {
		fmt.Printf("target found in slice\n\n")
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
