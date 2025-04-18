package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	newList := &ListNode{}
	current := newList

	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}

		current = current.Next
	}

	if list1 != nil {
		current.Next = list1
	} else if list2 != nil {
		current.Next = list2
	}

	return newList.Next
}

func main() {
	head1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	head2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	fmt.Print(mergeTwoLists(head1, head2))
}
