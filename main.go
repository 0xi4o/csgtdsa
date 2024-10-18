package main

import (
	"csgtdsa/hashtables"
	"csgtdsa/search"
	"csgtdsa/sort"
)

func main() {
	search.LinearSearch()
	search.BinarySearch()
	sort.BubbleSort()
	sort.SelectionSort()
	sort.InsertionSort()
	hashtables.RunIntersection()
}
