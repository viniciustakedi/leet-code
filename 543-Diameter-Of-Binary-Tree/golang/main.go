package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	var maxDiameter int

	var depth func(node *TreeNode) int
	depth = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		lDepth := depth(node.Left)
		rDepth := depth(node.Right)

		if lDepth+rDepth > maxDiameter {
			maxDiameter = lDepth + rDepth
		}

		return max(lDepth, rDepth) + 1
	}

	depth(root)
	return maxDiameter
}

func main() {
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val:   2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3},
	}

	fmt.Println(diameterOfBinaryTree(root))
	fmt.Println()

	root2 := &TreeNode{
		Val:  1,
		Left: &TreeNode{Val: 2},
	}

	fmt.Println(diameterOfBinaryTree(root2))
	fmt.Println()

	root3 := &TreeNode{
		Val: 1,
	}

	fmt.Println(diameterOfBinaryTree(root3))
	fmt.Println("DIASDKAKDAKSDSA")

	root4 := &TreeNode{
		Val: 2,
		Left: &TreeNode{
			Val:  3,
			Left: &TreeNode{Val: 1},
		},
	}

	fmt.Println(diameterOfBinaryTree(root4))
}
