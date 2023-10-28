package reverseLinkedList

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
		input    *linkedList.ListNode
		expected *linkedList.ListNode
	}{
		{0, linkedList.CreateList([]int{1, 2, 3, 4, 5}), linkedList.CreateList([]int{5, 4, 3, 2, 1})},
		{1, linkedList.CreateList([]int{1, 2}), linkedList.CreateList([]int{2, 1})},
		{2, linkedList.CreateList([]int{}), linkedList.CreateList([]int{})},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case #%v", tc.id), func(t *testing.T) {
			output := reverseList(tc.input)
			if !linkedList.Compare(output, tc.expected) {
				t.Errorf("\n case: %v\n output: %v\n expected: %v", linkedList.ToSlice(tc.input), linkedList.ToSlice(output), linkedList.ToSlice(tc.expected))
			}
		})
	}
}
