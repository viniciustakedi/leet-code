package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	hash := make(map[*ListNode]bool)

	for head != nil {
		if _, ok := hash[head]; ok {
			return true
		}

		hash[head] = true
		head = head.Next
	}

	return false
}

func main() {
	head := &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 0, Next: &ListNode{Val: -4, Next: nil}}}}
	fmt.Println(hasCycle(head))
}
