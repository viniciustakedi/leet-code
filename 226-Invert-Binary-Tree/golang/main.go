package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	response := invertRoot(root)
	return response
}

func invertRoot(root *TreeNode) *TreeNode {
	if root != nil {
		root.Left, root.Right = root.Right, root.Left
		invertRoot(root.Left)
		invertRoot(root.Right)
	}

	return root
}

func main() {
	root := &TreeNode{
		Val: 4,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 1},
			Right: &TreeNode{Val: 3},
		},
		Right: &TreeNode{
			Val:   7,
			Left:  &TreeNode{Val: 6},
			Right: &TreeNode{Val: 9},
		},
	}

	fmt.Print(invertTree(root))
}
