package search

import (
	"fmt"
)

func LinearSearch() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	result := Linear(nums, 5)
	if !result {
		fmt.Printf("target not found in slice")
	} else {
		fmt.Printf("target found in slice")
	}
}

func Linear(nums []int, target int) bool {
	for _, i := range nums {
		if i == target {
			return true
		}
	}

	return false
}
