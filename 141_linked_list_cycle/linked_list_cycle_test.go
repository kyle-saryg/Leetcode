package linkedListCycle

import (
	"fmt"
	"testing"

	"github.com/kyle-saryg/Leetcode/linkedList"
)

func TestHasCycle(t *testing.T) {
	testCases := []struct {
		id       int
		input    []int
		expected bool
	}{
		{0, []int{3, 2, 0, -4}, true},
		{0, []int{1, 2}, true},
		{0, []int{1}, false},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case #%v\n", tc.id), func(t *testing.T) {
			// Create linked list
			// Add the loop based on the 'expected' boolean
			ll := linkedList.CreateList(tc.input)
			if tc.expected {
				createLoop(ll)
			}

			result := hasCycle(ll)
			if !result {
				t.Errorf("\t\n--input: %v \n\t--expected: %v \n\t--result: %v\n", tc.input, tc.expected, result)
			}

		})
	}
}

func createLoop(head *linkedList.ListNode) {
	// Traverse to the last node of the list
	// Set next value to head
	//   -- Technically should set to a random node in the list
	//   -- However, setting it back to the head should suffice

	explorer := head

	for explorer.Next != nil {
		explorer = explorer.Next
	}

	explorer.Next = head
}
