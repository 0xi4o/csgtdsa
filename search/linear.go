// Chapter 1: Why Data Structures Matter
package search

import (
	"fmt"
)

func LinearSearch() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println("Linear Search:")
	result := Linear(nums, 5)
	if !result {
		fmt.Printf("target not found in slice\n\n")
	} else {
		fmt.Printf("target found in slice\n\n")
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
