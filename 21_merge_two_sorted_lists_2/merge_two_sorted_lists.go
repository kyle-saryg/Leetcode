package mergeTwoSortedLists2

import (
	. "github.com/kyle-saryg/Leetcode/linkedList"
)

/*
helper funcs:
  appendList()

*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	return nil
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
