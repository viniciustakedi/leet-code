package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var newList *ListNode

	for head != nil {
		nextNode := head.Next
		head.Next = newList
		newList = head
		head = nextNode
	}

	return newList
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	// head := &ListNode{}
	fmt.Print(reverseList(head))
}
