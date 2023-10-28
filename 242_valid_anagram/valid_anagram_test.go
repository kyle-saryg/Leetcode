package validAnagram

import (
	"fmt"
	"testing"
)

func TestValidAnagram(t *testing.T) {
	tc := []struct {
		id       int
		s        string
		t        string
		expected bool
	}{
		{0, "anagram", "nagaram", true},
		{1, "rat", "car", false},
		{2, "a", "ab", false},
	}

	for _, tc := range tc {
		t.Run(fmt.Sprintf("Test Case #%v", tc.id), func(t *testing.T) {
			output := isAnagram(tc.s, tc.t)
			if output != tc.expected {
				t.Errorf("\n input: %v %v\n expected: %v\n returned: %v", tc.s, tc.t, tc.expected, output)
			}
		})
	}
}
