package two_sum

import (
	"fmt"
	"reflect"
	"testing"
)

// func twoSum(nums []int, target int) []int
func TestTwoSum(t *testing.T) {
	testCases := []struct {
		id       int
		nums     []int
		target   int
		expected []int
	}{
		{0, []int{2, 7, 11, 15}, 9, []int{0, 1}},
		{1, []int{3, 2, 4}, 6, []int{1, 2}},
		{2, []int{3, 3}, 6, []int{0, 1}},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case #%v", tc.id), func(t *testing.T) {
			output := twoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(output, tc.expected) {
				t.Errorf("\n nums: %v\n target: %v\n output: %v\n expected: %v", tc.nums, tc.target, output, tc.expected)
			}
		})
	}
}
