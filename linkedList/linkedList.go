package linkedList

/*
TODO:
	toSlice(*ListNode) []int
*/

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// Returns head node with populated value
func InitList(val int) *ListNode {
	head := new(ListNode)
	head.Val = val

	head.Next = nil
	return head
}

// Creates a new node, appends to end of list
func (head *ListNode) Append(val int) {
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
func (head *ListNode) AppendSlice(vals []int) {
	for _, val := range vals {
		head.Append(val)
	}
}

func Compare(a *ListNode, b *ListNode) bool {
	if LinkedListLen(a) == 0 && LinkedListLen(b) == 0 {
		return true
	}

	if LinkedListLen(a) != LinkedListLen(b) {
		return false
	}

	// Iterate through either list
	// -- Compare values
	ptrA := a
	ptrB := b

	for ptrA.Next != nil {
		if ptrA.Val != ptrB.Val {
			return false
		}
		ptrA = ptrA.Next
		ptrB = ptrB.Next
	}

	return true
}

func CreateList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := InitList(vals[0])
	head.AppendSlice(vals[1:])
	return head
}

func createNode(val int) *ListNode {
	newNode := new(ListNode)
	newNode.Val = val
	newNode.Next = nil

	return newNode
}

func (head *ListNode) Display() {
	if head == nil {
		fmt.Printf("|%v|\n", nil)
		return
	}

	for head.Next != nil {
		fmt.Printf("|%v| -> ", head.Val)
		head = head.Next
	}
	fmt.Printf("|%v| -> |%v|", head.Val, head.Next)
	fmt.Printf("\n")
}

func LinkedListLen(head *ListNode) int {
	if head == nil {
		return 0
	}

	ptr := head
	ctr := 1
	for ptr.Next != nil {
		ptr = ptr.Next
		ctr++
	}

	return ctr
}

func ToSlice(head *ListNode) []int {
	var slice []int
	for head != nil {
		slice = append(slice, head.Val)
		head = head.Next
	}
	return slice
}
