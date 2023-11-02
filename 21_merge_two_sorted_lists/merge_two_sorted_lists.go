package mergeTwoSortedLists

import (
	"fmt"

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
	var prev1 *ListNode = nil
	/*
		Example:
		list 1: [1][2][4][nil]
					p1 c1
		list 2: [1][3][4][nil]
					c2 n2
		-- [1][1]
	*/

	for curr1 != nil && curr2 != nil {
		next2 := curr2.Next
		if curr1.Val >= curr2.Val {
			//Inserting curr2 into list1
			if prev1 != nil {
				prev1.Next = curr2
			}
			curr2.Next = curr1
			curr2 = next2
		} else {
			prev1 = curr1
			curr1 = curr1.Next
		}
	}

	fmt.Printf(" -- LINE 42: curr1: %v\tcurr2: %v\n", curr1, curr2)
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
