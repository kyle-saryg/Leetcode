package queue

import (
	"github.com/kyle-saryg/Leetcode/linkedList"
)

type Queue struct {
	queueNode *linkedList.ListNode
}

func (q *Queue) Add(val int) {
	// Adding to the end of the list
	q.queueNode.Append(val)
}

func InitStack() *Queue {
	return &Queue{queueNode: nil}
}
