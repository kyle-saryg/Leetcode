package mergeTwoSortedLists

import (
	. "github.com/kyle-saryg/Leetcode/linkedList"
)

/*
c1 starts at list1 head
c2 starts at list2 head

note: if c2 is nil, inserted all of list2 into list1, prevents seg fault
while c1 != nil and c2 != nil
  -- if c1 is GE c2
    -- insert c2 behind c1

if c2 != nil
  -- Append the rest of list2 to list1

helper funcs:
  appendList()
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	curr1 := list1
	curr2 := list2

	for curr1 != nil && curr2 != nil {
		next1 := curr1.Next
		next2 := curr2.Next
		if curr1.Val >= curr2.Val {
			//Inserting curr2 into list1
			curr1.Next = curr2
			curr2.Next = next1
			curr2 = next2
		} else {
			curr1 = next1
		}
	}

	if curr2 != nil {
		appendList()
	}

	// Inserting nodes from list2 into list1
	return list1
}

func appendList(destination *ListNode, target *ListNode) {

}
