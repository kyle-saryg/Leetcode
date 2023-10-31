package reverseLinkedList

import (
	. "github.com/kyle-saryg/Leetcode/linkedList"
)

/*
SINGLY LINKED LIST

Completely different approach in than algorithm 1 and 2.
Instead of reading and writing the values stored in each node,
redirecting pointers.

Algorithm: 3
Memory Complexity: O(1)

Time Complexity: O(n)
*/

/*
alg:
prev = nil
curr = head

while curr != nil
next = current.next
current.next = previous
previous = current
current = next

[1]->[2]->[3]->[4]->[nil]
c     n               p

[1]->[nil] [2]->[3]->[4]->[nil]
p           cn

[1]->[nil] [2]->[3]->[4]->[nil]
p          c     n

[2]->[1]->[nil] [3]->[4]->[nil]
c     p          n
*/

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode = nil
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	return prev
}
