package stack

import (
	"fmt"

	"github.com/kyle-saryg/Leetcode/linkedList"
)

/*
	Implemented using a linked list.
	Use a sentinel node, to demarcate the top of the list.
	Pop -> Insert at the start of the list
	Push -> Return the value at the start of the list
*/

type Stack struct {
	top *linkedList.ListNode
}

// Sentinel node points to the top of the stack
func InitStack() *Stack {

	return &Stack{nil}
}

func (s *Stack) Pop() (int, error) {
	if s.top == nil {
		return -1, fmt.Errorf("Stack is empty")
	}

	retVal := s.top.Val
	s.top = s.top.Next

	return retVal, nil
}

func (s *Stack) Push(nodeVal int) {
	// Inserting at the beginning of the linked list
	newNode := &linkedList.ListNode{Val: nodeVal, Next: s.top}

	// Setting new top to the beginning of the list
	s.top = newNode
}
