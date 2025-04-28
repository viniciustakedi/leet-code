package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	response, _ := remove(head, n)
	return response
}

func remove(head *ListNode, n int) (*ListNode, int) {
	if head == nil {
		return head, 0
	}

	ln, acc := remove(head.Next, n)

	head.Next = ln

	if acc+1 == n {
		return head.Next, acc + 1
	}

	return head, acc + 1
}

func main() {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	result := removeNthFromEnd(head, 2)
	var resultArray []int
	for result != nil {
		resultArray = append(resultArray, result.Val)
		result = result.Next
	}
	fmt.Println(resultArray)

	resultArray = []int{}

	head = &ListNode{Val: 1}
	result = removeNthFromEnd(head, 1)
	for result != nil {
		resultArray = append(resultArray, result.Val)
		result = result.Next
	}
	fmt.Println(resultArray)

	resultArray = []int{}

	head = &ListNode{Val: 1, Next: &ListNode{Val: 2}}
	result = removeNthFromEnd(head, 2)
	for result != nil {
		resultArray = append(resultArray, result.Val)
		result = result.Next
	}
	fmt.Println(resultArray)
}
