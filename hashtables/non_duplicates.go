// Chapter 8: Exercise 4
package hashtables

import (
	"fmt"
	"time"
)

func FindNonDuplicates() {
	input := "minimum"

	fmt.Println("NonDuplicates:")
	fmt.Printf("input: %s\n", input)
	now := time.Now()
	result := NonDuplicates(input)
	fmt.Printf("first non duplicate character: %c\n", result)
	fmt.Printf("time taken: %v\n\n", time.Since(now))
}

func NonDuplicates(input string) rune {
	var nonDuplicate rune
	hashtable := make(map[rune]int)

	for _, val := range input {
		hashtable[val] += 1
	}

	for k, v := range hashtable {
		if v == 1 {
			nonDuplicate = rune(k)
			break
		}
	}

	return nonDuplicate
}
