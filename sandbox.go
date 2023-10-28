package main

import (
	"github.com/kyle-saryg/Leetcode/linkedList"
)

func main() {
	head := linkedList.InitList(3)
	head.AppendSlice([]int{1, 6, 3, 8, 0, 2, 4, 8, 4, 9})

	head.Display()
}
