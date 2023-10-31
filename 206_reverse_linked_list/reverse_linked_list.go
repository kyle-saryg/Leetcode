package reverseLinkedList

import (
	. "github.com/kyle-saryg/Leetcode/linkedList"
)

/*
SINGLY LINKED LIST

Algorithm: 1
-- Find index of last node

-- Iterate through LL backwards
  - Copying values into a slice or array
  - Slice is now in reverse order from LL

-- Iterate through LL
  - Copying values from slice into each node
  - NOTE: Use an indexing variable to synchronize slice and LL nodes

Notes:
  - Pretty bad Memory - Complexity
  - Could make algorithm in-place
  - Replace slice with a single int (would add more traversal)
  - Sacrifice memory, for speed

Memory Complexity: O(n)

Time Complexity: O(n^2 + n + n) => O(n^2)
  - Finding Max index = O(n)
  - Traversing LL Backwards = O(n^2)
  - Traversing LL = O(n)
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

	// Traverse through LL backwards
	// Copying values into slice
	values := []int{}
	for i := maxIndex; i >= 0; i-- {
		ptr = moveTo(head, i)
		values = append(values, ptr.Val)
	}

	// Traverse LL copying values from slice to each corresponding node
	ptr = head
	for i := 0; i <= maxIndex; i++ {
		ptr.Val = values[i]
		ptr = ptr.Next
	}

	return head
}

/*
Usage:
-- Setting 'index' to a value past the list WILL SEGFAULT

TODO:
-- Error check if 'ptr' moves past the end of list
*/
// SEGFAULTS DEBUG
func moveTo(ptr *ListNode, index int) *ListNode {
	tmp := 0
	// TODO: Check if < or <=

	/*
		index = 4

		[] [] [] [] []
		0  1  2  3  4
	*/

	for tmp < index {
		ptr = ptr.Next
		tmp++
	}

	return ptr
}
