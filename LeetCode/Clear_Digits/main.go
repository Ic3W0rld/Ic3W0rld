package main

import (
	"fmt"
	"unicode"
)

func clearDigits(s string) string {
	for {
		// Flag to check if any digit was found and processed
		processed := false

		// Iterate over the string to find the first digit
		for i := 0; i < len(s); i++ {
			if unicode.IsDigit(rune(s[i])) {
				// Check if the character to the left is a non-digit
				if i > 0 && !unicode.IsDigit(rune(s[i-1])) {
					// Remove the digit and the non-digit to its left
					s = s[:i-1] + s[i+1:]
					processed = true
					break // Restart the loop after modification
				} else if i == 0 {
					// If the digit is at the beginning, just remove the digit
					s = s[i+1:]
					processed = true
					break // Restart the loop after modification
				}
			}
		}

		// If no digit was processed in this iteration, exit the loop
		if !processed {
			break
		}
	}

	return s
}

func main() {
	s1 := "abc"
	fmt.Println(clearDigits(s1))

	s2 := "cb34"
	fmt.Println(clearDigits(s2))

}
