package validPalindrome

import (
	"strings"
	"unicode"
)

/*
Algorithm:
  1. Filter string for non-alphanumeric characters (commas, periods, question marks, etc.)
    1.1 Use unicode library
      -- unicode.IsLetter()
      -- unicode.IsNumber()
      -- Create a helper function alphanumericFilter()
	1.2 Convert all characters to lowercase

  2. Verify palindrome
    -- Place p1 at the start of the string
	-- Place p2 at the end of the string
	  -- Check if they have the same character
	-- Increment p1
	-- Decrement p2

  3. Repeat 2 intil p1 and p2 are pointing at the same index or p1 becomes greater than p2
    -- p1 <= p2
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
