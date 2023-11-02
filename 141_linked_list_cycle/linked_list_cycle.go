package linkedListCycle

import (
	. "github.com/kyle-saryg/Leetcode/linkedList"
)

/*
Note: Two algorithms in mind, want to write both.
 1. Runner technique for detecting loops in linked lists.
 2. Create a hashmap, with the memory location of each visited node as a key, and a bool as value (the value doesn't really matter too much)
   -- Traverse the linked list, adding each node to the hash
   -- If the memory location of the node exists in the hash
     -- A loop within the linked list exists


Algorithm 1:
 -- fastRunner starts at the head of linked list
 -- slowRunner starts at the head of linked list

 -- While fastRunner hasn't reach the end of the list
   -- fastRunner moves two nodes
   -- slowRunner moves one node

   -- If fastRunner and slowrunner point at the same node
     -- Linked list has a loop

 -- Reached the end of the list
   -- Linked List does NOT have a loop
*/

func hasCycle(head *ListNode) bool {
	/*
		TODO:
			-- May not need both conditions in part 1
			-- Edgecase: Empty list
			?- Edgecase: Single list node looping back on itself
	*/
	fastRunner := head
	slowRunner := head

	// 1. May not need both conditions
	for fastRunner.Next != nil && fastRunner != nil {
		// Traverse fastRunner two nodes
		fastRunner = fastRunner.Next
		fastRunner = fastRunner.Next
		// Traverse slowRunner one node
		slowRunner = slowRunner.Next

		// Are they pointing at the same node?
		if fastRunner == slowRunner {
			// loop exists
			return true
		}
	}

	// No loop
	return false
}
