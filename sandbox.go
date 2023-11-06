package main

import (
	"fmt"

	"github.com/kyle-saryg/Leetcode/linkedList"
	"github.com/kyle-saryg/Leetcode/stack"
)

func main() {
	testStack()
}

func testStack() {
	stack := stack.InitStack()

	for i := 0; i <= 10; i++ {
		fmt.Printf("Pushing -> %v\n", i)
		stack.Push(i)
	}

	for retVal, err := stack.Pop(); err == nil; retVal, err = stack.Pop() {
		fmt.Printf("Popping -> %v\n", retVal)
	}

	_, err := stack.Pop()
	fmt.Printf("err -> %v\n", err)
}

func testLinkedList() {
	headA := linkedList.CreateList([]int{1, 5, 2, 7, 2, 8})
	headB := linkedList.CreateList([]int{6, 3, 8, 93, 4})

	headA.Display()
	fmt.Printf("Length of list: %v\n", linkedList.LinkedListLen(headA))
	headB.Display()
	fmt.Printf("Length of list: %v\n", linkedList.LinkedListLen(headB))

	if linkedList.Compare(headA, headB) {
		fmt.Println("EQUIVALENT")
	} else {
		fmt.Println("NOT EQUIVALENT")
	}

	fmt.Printf("%v\n", linkedList.ToSlice(headA))
	fmt.Printf("%v\n", linkedList.ToSlice(headB))
}
