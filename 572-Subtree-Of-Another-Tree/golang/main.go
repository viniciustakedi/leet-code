package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isSubtree(root *TreeNode, subRoot *TreeNode) bool {
	if root == nil && subRoot == nil {
		return true
	}

	if root == nil || subRoot == nil {
		return false
	}

	if root.Val == subRoot.Val && isSameTree(root, subRoot) {
		return true
	}

	return isSubtree(root.Left, subRoot) || isSubtree(root.Right, subRoot)
}

func isSameTree(q *TreeNode, p *TreeNode) bool {
	if p == q {
		return true
	}

	if p == nil || q == nil || p.Val != q.Val {
		return false
	}

	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}

func main() {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:   4,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 2},
		},
		Right: &TreeNode{Val: 5},
	}

	subRoot := &TreeNode{
		Val:   4,
		Left:  &TreeNode{Val: 1},
		Right: &TreeNode{Val: 2},
	}

	fmt.Println(isSubtree(root, subRoot))

	root2 := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val:  4,
			Left: &TreeNode{Val: 1},
			Right: &TreeNode{
				Val:  2,
				Left: &TreeNode{Val: 0},
			},
		},
		Right: &TreeNode{Val: 5},
	}

	fmt.Println(isSubtree(root2, subRoot))
}
