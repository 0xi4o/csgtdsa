// Chapter 8: Exercise 2
package hashtables

import (
	"fmt"
	"time"
)

func FindFirstDuplicate() {
	input := "abcdcef"

	fmt.Println("FirstDuplicate:")
	fmt.Printf("input: %s\n", input)
	now := time.Now()
	result := FirstDuplicate(input)
	fmt.Printf("first duplicate character: %c\n", result)
	fmt.Printf("time taken: %v\n\n", time.Since(now))
}

func FirstDuplicate(input string) rune {
	var duplicate rune
	hashtable := make(map[rune]int)

	for _, val := range input {
		hashtable[val] += 1
	}

	fmt.Printf("%v\n", hashtable)

	for k, v := range hashtable {
		if v > 1 {
			duplicate = rune(k)
			break
		}
	}

	return duplicate
}
