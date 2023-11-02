package mergeTwoSortedLists

import (
	. "github.com/kyle-saryg/Leetcode/linkedList"
)

/*
c1 starts at list1 head
c2 starts at list2 head

note: if c2 is nil, inserted all of list2 into list1, prevents seg fault
while c1 != nil and c2 != nil
  -- n1 = c1.Next
  -- n2 = n2.Next

  -- if c1 is LE c2
    -- insert c2 ahead c1
	-- c1 = c2 // Inserted c2 into list 1
	-- c2 = next2 // Moving c2 back into list 2

if c2 != nil
  -- Append the rest of list2 to list1

helper funcs:
  appendList()

*/
/*
Example:
	c1 = list 1 current
	n2 = leading c1 // Used to connect list 1 when inserting c2
	c2 = list 2 current
	n2 = leading c2 // Used to move c2 back to list 2 (Post - Insertion)


	list 1: [1][2][4][nil]
	list 2: [1][3][4][nil]
	--
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// Inserting nodes from list2 into list1
	return list1
}

// 'dest' is the end of list1
// 'target' is where 'curr2' leaves off at
func appendList(dest *ListNode, target *ListNode) {
	for target != nil {
		dest.Next = target
		dest = target
		target = target.Next
	}
}
