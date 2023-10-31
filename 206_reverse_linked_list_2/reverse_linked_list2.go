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
   -- ptr2 moves backard one node
   -- swap

 - Repeat Step A and B until:
   -- ptr1 and ptr2 point at the same node
      OR
   -- ptr1 is one node behind ptr2

Space Complexity: O(1)
 - In-place algorithm
 - Unlike first algorithm
   -- slice which grows with the size of the LL
 - No additional memory is allocated
*/

func reverseList(head *ListNode) *ListNode {
	return nil
}
