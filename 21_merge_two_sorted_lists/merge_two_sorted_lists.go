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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil && list2 == nil {
		return nil
	}

	curr1 := list1
	curr2 := list2
	var prev1 *ListNode = nil
	next2 := curr2.Next
	/*
		Example:
		list 1: [1][2][4][nil]
		        c1
		list 2: [1][3][4][nil]
				c2  n2
				p1
		--
	*/

	for curr1 != nil && next2 != nil {

		if curr1.Val >= curr2.Val {
			//Inserting curr2 into list1
			if prev1 == nil {
				// Inserting behind the head of list1
				// Re-assigning head
				list1 = curr2
			}
			if prev1 != nil {
				prev1.Next = curr2
			}
			prev1 = curr2
			curr2.Next = curr1
			curr2 = next2
		} else {
			prev1 = curr1
			curr1 = curr1.Next
		}

		next2 = curr2.Next
	}

	if curr2 != nil {
		appendList(curr1, curr2)
	}

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
