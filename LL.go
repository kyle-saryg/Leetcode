package linkedList

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
func (*ListNode) append(val int) {

}
