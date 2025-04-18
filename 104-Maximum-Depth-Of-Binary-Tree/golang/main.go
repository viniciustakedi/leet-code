package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func maxDepth(root *TreeNode) int {
	return countNodes(root)
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	amountLeft := 0
	amountRight := 0

	if root.Left != nil {
		amountLeft += countNodes(root.Left)
	}

	if root.Right != nil {
		amountRight += countNodes(root.Right)
	}

	amount := 1

	if amountLeft > amountRight {
		amount += amountLeft
	} else {
		amount += amountRight
	}

	return amount
}

func main() {
	root := &TreeNode{
		Val:  3,
		Left: &TreeNode{Val: 9},
		Right: &TreeNode{
			Val:   20,
			Left:  &TreeNode{Val: 15},
			Right: &TreeNode{Val: 7},
		},
	}

	// root := &TreeNode{
	// 	Val:   1,
	// 	Left:  nil,
	// 	Right: &TreeNode{Val: 2},
	// }

	fmt.Print(maxDepth(root))
}
