package main

import (
	"fmt"

	"github.com/kyle-saryg/Leetcode/linkedList"
)

func main() {
	headA := linkedList.CreateList([]int{})
	headB := linkedList.CreateList([]int{})

	headA.Display()
	fmt.Printf("Length of list: %v\n", linkedList.LinkedListLen(headA))
	headB.Display()
	fmt.Printf("Length of list: %v\n", linkedList.LinkedListLen(headB))

	if linkedList.Compare(headA, headB) {
		fmt.Println("EQUIVALENT")
	} else {
		fmt.Println("NOT EQUIVALENT")
	}

}
