package linkedList

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Returns head node with populated value
func initList(val int) *ListNode {
	head := new(ListNode)
	head.Val = val

	head.Next = nil
	return head
}

// Creates a new node, appends to end of list
func (head *ListNode) append(val int) {
	// Traversing Linked List (stops at last node)
	var nodePtr *ListNode = head
	for nodePtr.Next != nil {
		nodePtr = nodePtr.Next
	}

	// New node creation and insertion
	newNode := createNode(val)
	nodePtr.Next = newNode
}

// Slice is appended in-order
func (head *ListNode) appendSlice(val []int) {

	return nil
}

func createNode(val int) *ListNode {
	newNode := new(ListNode)
	newNode.Val = val
	newNode.Next = nil

	return newNode
}

func (head *ListNode) display() {
	for head != nil {
		fmt.Printf("<%v> -> ", head.Val)
		head = head.Next
	}
	fmt.Printf("\n")
}
