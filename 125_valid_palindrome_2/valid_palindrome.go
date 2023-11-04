package validPalindrome2

import (
	"strings"
	"unicode"
)

/*
Algorithm:
  -- Place p1 at the start of the string
  -- Place p2 at the end of the string


  1. Verify palindrome
    -- if p1 is NOT alphanumeric
	  -- Increment p1
	  -- pass
	-- if p2 is NOT alphanumeric
	  -- Decrement p1
	  -- pass

	-- Ensure lowercase of p1 and p2 are the same character

	-- Increment p1
	-- Decrement p2

  2. Repeat 2 intil p1 and p2 are pointing at the same index or p1 becomes greater than p2
    -- p1 <= p2

Time - Complexity:

Space - Complexity:
*/

func isPalindrome(s string) bool {
	//1. Filter input string
	filteredString := alphanumericFilter(s)
	input := strings.ToLower(filteredString)

	trailIndex := 0
	leadIndex := len(input) - 1

	//2 & 3, verifying palindrome
	for trailIndex <= leadIndex {
		if input[trailIndex] != input[leadIndex] {
			return false
		}
		trailIndex++
		leadIndex--
	}

	return true
}

func alphanumericFilter(input string) string {
	filteredString := ""

	for _, char := range input {
		if unicode.IsLetter(char) || unicode.IsNumber(char) {
			filteredString = filteredString + string(char)
		}
	}

	return filteredString
}
