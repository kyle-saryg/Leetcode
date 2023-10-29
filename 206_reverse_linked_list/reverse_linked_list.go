package reverseLinkedList

import (
	. "github.com/kyle-saryg/Leetcode/linkedList"
)

/*
SINGLY LINKED LIST

Algorithm:

	 -- Find index of last node

	 -- Iterate through LL backwards
		 -- Copying values into a slice or array
		 -- Slice is now in reverse order from LL

	 -- Iterate through LL
	 	 -- Copying values from slice into each node
		 -- NOTE: Use an indexing variable to synchronize slice and LL nodes
*/
func reverseList(head *ListNode) *ListNode {
	// Empty list
	if head == nil {
		return head
	}

	// First Node is index 0
	// Last Node is 'maxIndex - 1'
	maxIndex := 0
	ptr := head
	for ptr.Next != nil {
		maxIndex++
		ptr = ptr.Next
	}

	// Iterate through LL backwards
	// Copying values into slice
	values := []int
	i := maxIndex
	for i >= 0 {
		ptr = moveTo(ptr, i)
	}

	return nil
}

/*
Usage:

	-- Setting 'index' to a value past the list WILL SEGFAULT

TODO:

	-- Check if 'ptr' moves past the end of list
*/
func moveTo(ptr *ListNode, index int) *ListNode {

}
