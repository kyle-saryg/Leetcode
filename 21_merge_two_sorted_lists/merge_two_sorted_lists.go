package mergeTwoSortedLists

import (
	. "github.com/kyle-saryg/Leetcode/linkedList"
)

/*
algorithm:
 -- curr1 starts at head of list1
 -- next1 is one node ahead of curr1
 -- curr2 starts at head of list2
 -- next2 is one node ahead of curr2

 -- Compare values of curr1 and curr2
 -- Example: curr1 is less than curr2
   -- curr1.next = curr2
   -- curr2.next = next1
   -- curr1 = next1
   -- next1 = next1.next
*/

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	return nil
}
