package hashtables

import (
	"fmt"
	"time"
)

func RunIntersection() {
	nums1 := []int{1, 2, 3, 4, 5}
	nums2 := []int{0, 2, 4, 6, 8}

	fmt.Println("Intersection:")
	fmt.Printf("first slice: %v\n", nums1)
	fmt.Printf("second slice: %v\n", nums2)
	now := time.Now()
	result := Intersection(nums1, nums2)
	fmt.Printf("intersection slice: %v\n", result)
	fmt.Printf("time taken: %v\n\n", time.Since(now))
}

func Intersection(s1, s2 []int) []int {
	var intersection []int
	hashtable := make(map[int]bool)

	for _, val := range s1 {
		hashtable[val] = true
	}

	for _, val := range s2 {
		if hashtable[val] {
			intersection = append(intersection, val)
		}
	}

	return intersection
}
