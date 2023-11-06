package queue

import (
	"github.com/kyle-saryg/Leetcode/linkedList"
)

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

// Returns the next node in the queue AND deletes
func (q *Queue) QueueUp() int {
}

func InitStack() *Queue {
	return &Queue{firstNode: nil, lastNode: nil}
}
