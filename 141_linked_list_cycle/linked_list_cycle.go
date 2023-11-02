package linkedListCycle

import (
	. "github.com/kyle-saryg/Leetcode/linkedList"
)

/*
Note: Two algorithms in mind, want to write both.
 -- Runner technique for detecting loops in linked lists.
 -- Create a hashmap, with the memory location of each visited node as a key, and a bool as value (the value doesn't really matter too much)
   -- Traverse the linked list, adding each node to the hash
   -- If the memory location of the node exists in the hash
     -- A loop within the linked list exists


Algorithm
*/

func hasCycle(head *ListNode) bool {}
