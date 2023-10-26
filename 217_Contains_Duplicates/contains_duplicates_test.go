package contains_duplicates

import (
	"fmt"
	"testing"
)

func TestContainsDuplicates(t *testing.T) {
	testCases := []struct {
		id       int
		nums     []int
		expected bool
	}{
		{0, []int{1, 2, 3, 1}, true},
		{1, []int{1, 2, 3, 4}, false},
		{2, []int{1, 1, 1, 3, 3, 4, 3, 2, 4, 2}, true},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case #%v", tc.id), func(t *testing.T) {
			output := containsDuplicates(tc.nums)
			if tc.expected != output {
				t.Errorf("\n input: %v\n expected: %v\n returned: %v", tc.nums, tc.expected, output)
			}
		})
	}
}
