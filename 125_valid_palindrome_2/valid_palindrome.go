package validPalindrome2

import (
	"unicode"
)

/*
Algorithm:
  -- Place p1 at the start of the string
  -- Place p2 at the end of the string


  1. Verify palindrome
    1.1 If p1 or p2 is NOT alphanumeric
	  -- Move the non-alphanumeric pointer
	  -- pass

	1.2 If lowercase of p1 and p2 are not equal
	  -- Not a palindrome

	-- Increment p1
	-- Decrement p2

  2. Repeat 2 intil p1 and p2 are pointing at the same index or p1 becomes greater than p2
    -- p1 <= p2

Time - Complexity:

Space - Complexity:
*/

func isPalindrome(s string) bool {
	trailing := 0
	leading := len(s) - 1

	//1. Verifying palindrome
	// Could also be < ?? check later
	for trailing <= leading {
		trailChar := rune(s[trailing])
		leadChar := rune(s[leading])

		// 1.1 Ensuring iterators are alphanumeric
		if !unicode.IsLetter(trailChar) || !unicode.IsDigit(trailChar) {
			trailing++
			continue
		}
		if !unicode.IsLetter(leadChar) || !unicode.IsDigit(leadChar) {
			leading--
			continue
		}

		// 1.2 Verifying equivalence
		if unicode.ToLower(trailChar) != unicode.ToLower(leadChar) {
			return false
		}

		trailing++
		leading--
	}

	// Verified entire string
	return true
}
