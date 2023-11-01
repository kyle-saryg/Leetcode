package mergeTwoSortedLists

import (
	"fmt"
	"testing"

	"github.com/kyle-saryg/Leetcode/linkedList"
)

func TestMergeTwoLists(t *testing.T) {
	testCases := []struct {
		id       int
		list1    []int
		list2    []int
		expected []int
	}{
		{0, []int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4}},
		{1, []int{}, []int{}, []int{}},
		{2, []int{}, []int{0}, []int{0}},
	}
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Test Case #%v", tc.id), func(t *testing.T) {
			output := mergeTwoLists(linkedList.CreateList(tc.list1), linkedList.CreateList(tc.list2))

			if !linkedList.Compare(output, linkedList.CreateList(tc.expected)) {
				// Fill out t.errorf
			}

		})
	}
}
