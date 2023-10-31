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

 - Repeat Step A and B until:
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
	// ptr1 starts at the beginning of the list, iterates forward
	ptr1 := head
	// ptr2 starts at the end of the list, iterates back
	ptr2 := head
	// Used when swapping
	var tmp int

	// Moving ptr2 to the end of the list
	for ptr2.Next != nil {
		ptr2 = ptr2.Next
	}

	return head
}
