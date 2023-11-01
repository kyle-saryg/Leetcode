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

	return nil
}
