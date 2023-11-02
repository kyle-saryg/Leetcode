package linkedListCycle2

import (
	"fmt"

	. "github.com/kyle-saryg/Leetcode/linkedList"
)

/*
Note: Two algorithms in mind, want to write both.
 1. Runner technique for detecting loops in linked lists.
 2. Create a hashmap, with the memory location of each visited node as a key, and a bool as value (the value doesn't really matter too much)
   -- Traverse the linked list, adding each node to the hash
   -- If the memory location of the node exists in the hash
     -- A loop within the linked list exists

Algorithm 2: (Worse space complexity that algorithm 1, same time complexity)
 -- explorer starts at the head of the list

 1. Create hashmap of hex string memory address keys, bool values
   -- fmt.Sprintf() => %p converts memory addresses into hexadecimal strings

 2. While explorer hasn't reached the end of the list
   -- Check if memory location is in the hashmap
     -- If memory location exists
	 -- Loop is present

   -- Add memory location of node into hashmap
   -- Move to the next node

 3. Traversed entire list
   -- No loop is present

Time - Complexity: O(n)
 -- Single traversal of the linked list

Space Complexity: O(n)
 -- Hash map grows with the size of the linked list
*/

func hasCycle(head *ListNode) bool {
	if head == nil {
		return false
	}

	hash := make(map[string]bool)
	explorer := head

	for explorer != nil {
		key := fmt.Sprintf("%p", explorer)

		// Check if value is in the list
		_, exists := hash[key]
		if exists {
			return true
		}

		// Insert into hashmap
		// The value doesn't really matter, we are just using the hash to
		// check if the key has been inserted
		hash[key] = true

		// moving to the next node
		explorer = explorer.Next
	}

	return false
}
