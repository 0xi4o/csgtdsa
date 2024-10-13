// Chapter 6: Optimizing for Optimistic Scenarios
package sort

import (
	"fmt"
	"time"
)

func InsertionSort() {
	nums := []int{73, 28, 91, 45, 12, 67, 36, 82, 19, 54, 3, 96, 40, 8, 61, 25, 79, 50, 14, 88}

	fmt.Println("Insertion Sort:")
	fmt.Printf("original slice: %v\n", nums)
	now := time.Now()
	result := Insertion(nums)
	fmt.Printf("sorted slice: %v\n", result)
	fmt.Printf("time taken: %v\n\n", time.Since(now))
}

func Insertion(nums []int) []int {
	i := 1
	for i < len(nums) {
		temp := nums[i]
		position := i - 1

		for position >= 0 {
			if nums[position] > temp {
				nums[position+1] = nums[position]
				position--
			} else {
				break
			}
			nums[position+1] = temp
		}
		i++
	}
	return nums
}
