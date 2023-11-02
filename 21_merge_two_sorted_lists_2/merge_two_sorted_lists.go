package mergeTwoSortedLists2

import (
	. "github.com/kyle-saryg/Leetcode/linkedList"
)

/*
Algorithm:
	1. Select head of list
	2. Traverse both lists
	  -- Compare values, insert lowest value node into new list
	3. Once the end of either list is reached
	  -- append the rest of the other list

helper funcs:
  appendList()

*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	var newHead *ListNode

	// 1. Selecting head of new list
	if list1.Val <= list2.Val {
		newHead = list1
		list1 = list1.Next
	} else {
		newHead = list2
		list2 = list2.Next
	}

	ptr := newHead

	// 2. Inserting nodes to new list
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			ptr.Next = list1
			list1 = list1.Next
		} else {
			ptr.Next = list2
			list2 = list2.Next
		}
		ptr = ptr.Next
	}

	// 3. appending the rest of the nodes
	if list1 == nil {
		appendList(ptr, list2)
	} else {
		appendList(ptr, list1)
	}

	return newHead
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
