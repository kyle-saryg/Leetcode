package validPalindrome

import (
	"fmt"
	"testing"
)

// func isPalindrome(s string) bool {}

func TestIsPalindrome(t *testing.T) {
	testCases := []struct {
		id       int
		input    string
		expected bool
	}{
		{0, "A man, a plan, a canal: Panama", true},
		{1, "race a car", false},
		{2, " ", true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case #%v", tc.id), func(t *testing.T) {
			output := isPalindrome(tc.input)

			if output != tc.expected {
				t.Errorf("\t -- input: %v\n\t -- expected %v\n\t -- output %v\n", tc.input, tc.expected, output)
			}
		})
	}
}
