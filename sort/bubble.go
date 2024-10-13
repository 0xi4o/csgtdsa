// Chapter 4: Speeding Up Your Code with Big O
package sort

import (
	"fmt"
	"time"
)

func BubbleSort() {
	nums := []int{73, 28, 91, 45, 12, 67, 36, 82, 19, 54, 3, 96, 40, 8, 61, 25, 79, 50, 14, 88}

	fmt.Println("Bubble Sort:")
	fmt.Printf("original slice: %v\n", nums)
	now := time.Now()
	result := Bubble(nums)
	fmt.Printf("sorted slice: %v\n", result)
	fmt.Printf("time taken: %v\n\n", time.Since(now))
}

func Bubble(nums []int) []int {
	unsortedEndIndex := len(nums) - 1
	isSorted := false

	for !isSorted {
		isSorted = true
		for i := range unsortedEndIndex {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				isSorted = false
			}
		}
		unsortedEndIndex -= 1
	}

	return nums
}
