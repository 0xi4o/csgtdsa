package sort

import "fmt"

func BubbleSort() {
	nums := []int{73, 28, 91, 45, 12, 67, 36, 82, 19, 54, 3, 96, 40, 8, 61, 25, 79, 50, 14, 88}

	fmt.Printf("original slice: %v\n", nums)
	result := Bubble(nums)
	fmt.Printf("sorted slice: %v\n", result)
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
