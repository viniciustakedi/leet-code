package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// func isSameTree(p *TreeNode, q *TreeNode) bool {
// 	if p == nil && q == nil {
// 		return true
// 	}

// 	if p == nil && q != nil {
// 		return false
// 	}

// 	if q == nil && p != nil {
// 		return false
// 	}

// 	if p.Val != q.Val {
// 		return false
// 	}

// 	if p.Left == nil && q.Left != nil {
// 		return false
// 	} else if q.Left == nil && p.Left != nil {
// 		return false
// 	}

// 	if p.Right == nil && q.Right != nil {
// 		return false
// 	} else if q.Right == nil && p.Right != nil {
// 		return false
// 	}

// 	rightResponse := true
// 	leftResponse := true

// 	if p.Right != nil && q.Right != nil {
// 		rightResponse = isSameTree(p.Right, q.Right)
// 	}

// 	if p.Left != nil && q.Left != nil {
// 		leftResponse = isSameTree(p.Left, q.Left)
// 	}

// 	if !leftResponse || !rightResponse {
// 		return false
// 	}

// 	return true
// }

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == q {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	p := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 2},
		Right: &TreeNode{Val: 1},
	}

	q := &TreeNode{
		Val:   1,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}

	fmt.Println(isSameTree(p, q), "ORIUAUERAIEIUS")

	p2 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}

	q2 := &TreeNode{
		Val:   1,
		Right: &TreeNode{Val: 2},
	}

	fmt.Println(isSameTree(p2, q2))
}
