package reverseLinkedList2

import (
	"fmt"
	"testing"

	"github.com/kyle-saryg/Leetcode/linkedList"
)

func TestReverseList(t *testing.T) {
	/*
		Input: head = [1,2,3,4,5]
		Output: [5,4,3,2,1]

		Input: head = [1,2]
		Output: [2,1]

		Input: head = []
		Output: []
	*/

	testCases := []struct {
		id       int
		input    []int
		expected []int
	}{
		{0, []int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{1, []int{1, 2}, []int{2, 1}},
		{2, []int{}, []int{}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case #%v", tc.id), func(t *testing.T) {
			output := reverseList(linkedList.CreateList(tc.input))
			if !linkedList.Compare(output, linkedList.CreateList(tc.expected)) {
				t.Errorf("\n case: %v\n output: %v\n expected: %v", tc.input, linkedList.ToSlice(output), tc.expected)
			}
		})
	}
}
