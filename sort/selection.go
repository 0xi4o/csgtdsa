// Chapter 5: Optimizing Code With and Without Big O
package sort

import (
	"fmt"
	"time"
)

func SelectionSort() {
	nums := []int{73, 28, 91, 45, 12, 67, 36, 82, 19, 54, 3, 96, 40, 8, 61, 25, 79, 50, 14, 88}

	fmt.Println("Selection Sort:")
	fmt.Printf("original slice: %v\n", nums)
	now := time.Now()
	result := Selection(nums)
	fmt.Printf("sorted slice: %v\n", result)
	fmt.Printf("time taken: %v\n\n", time.Since(now))
}

func Selection(nums []int) []int {
	for i := range nums {
		lowest := i
		j := i + 1
		for j < len(nums) {
			if nums[j] < nums[lowest] {
				lowest = j
			}
			j++
		}
		if lowest != i {
			nums[i], nums[lowest] = nums[lowest], nums[i]
		}
	}

	return nums
}
