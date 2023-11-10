package queue

import (
	"fmt"

	"github.com/kyle-saryg/Leetcode/linkedList"
)

/*
Implemented a queue using a linked list
*/

type Queue struct {
	firstNode *linkedList.ListNode
	lastNode  *linkedList.ListNode
}

func (q *Queue) Add(val int) {
	newNode := &linkedList.ListNode{Val: val, Next: nil}

	// Error checking, prevents seg-faults when inserting into an empty queue
	if q.firstNode == nil {
		q.firstNode = newNode
	}
	if q.lastNode == nil {
		q.lastNode = newNode
	}

	// Inserting into the end
	q.lastNode.Next = newNode
	// Updating last node
	q.lastNode = newNode
}

// Returns the next node in the queue AND deletes it (removing all references to the node)
func (q *Queue) QueueUp() (int, error) {
	/*
		[front] -> [] -> [] -> [] -> [end]
		q.FN                          q.LN
	*/
	if q.firstNode == nil {
		return -1, fmt.Errorf("Queue is empty")
	}

	retVal := q.firstNode.Val
	q.firstNode = q.firstNode.Next

	return retVal, nil
}

func InitQueue() *Queue {
	return &Queue{firstNode: nil, lastNode: nil}
}
