package reverseLinkedList2

import (
	. "github.com/kyle-saryg/Leetcode/linkedList"
)

/*
Algorithm: (Sacrificing time complexity for more efficient memory complexity)
 - Step A:
   -- Move ptr1 to first node
   -- Move ptr2 to last node
   -- Swap values

 - Step B:
   -- ptr1 moves forward one node
   -- ptr2 moves back one node
   -- Swap values

  - Step C: Repeat Step A and B until
   -- ptr1 and ptr2 point at the same node
      OR
   -- ptr1 is one node behind ptr2

Space Complexity: O(1)
 - In-place algorithm
 - Unlike first algorithm
   -- slice which grows with the size of the LL
 - No additional memory is allocated

Time Complexity: O(n^2) ?? CHECK ONCE DONE CODING
*/

func reverseList(head *ListNode) *ListNode {
	// Empty List
	if head == nil {
		return nil
	}

	// ptr1 starts at the beginning of the list, iterates forward
	ptr1 := head
	// ptr2 starts at the end of the list, iterates back
	ptr2 := head
	// Index is initialized to the max index of the list, then decremented during swappage
	index := 0

	// Moving ptr2 to the end of the list
	for ptr2.Next != nil {
		ptr2 = ptr2.Next
		index++
	}

	// Infinite loop, breaks on condition C
	for {
		// Swap values
		swapValues(ptr1, ptr2)
		/*
		   Post-swap:
		    -- ptr1 is one node behind ptr2
		    -- ptr1 and ptr2 are pointing at the same node
		*/
		if ptr1.Next == ptr2 || ptr1 == ptr2 {
			break
		}
		// Move ptr1 forward
		ptr1 = ptr1.Next
		// Move ptr2 backward
		ptr2 = moveTo(head, index)
		index--
	}

	return head
}

func swapValues(node1 *ListNode, node2 *ListNode) {
	tmp := node1.Val
	node1.Val = node2.Val
	node2.Val = tmp
}

/*
Caution:
 1. if index is past the last node, WILL SEGFAULT

TODO:
  - Fix 1
*/
func moveTo(head *ListNode, index int) *ListNode {
	// head is at index 0
	for ctr := 0; ctr < index; ctr++ {
		head = head.Next
	}

	return head
}
