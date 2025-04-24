package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, balanced := height(root)
	return balanced
}

func height(node *TreeNode) (int, bool) {
	if node == nil {
		return 0, true
	}

	leftHeight, leftBalanced := height(node.Left)
	rightHeight, rightBalanced := height(node.Right)

	if !leftBalanced || !rightBalanced || abs(leftHeight-rightHeight) > 1 {
		return 0, false
	}

	return max(leftHeight, rightHeight) + 1, true
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func main() {

	// Case 1: [3,9,20,null,null,15,7]
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	// Case 2: [1,2,2,3,3,null,null,4,4]
	rootTwo := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
				Left: &TreeNode{
					Val: 4,
				},
				Right: &TreeNode{
					Val: 4,
				},
			},
			Right: &TreeNode{
				Val: 3,
			},
		},
		Right: &TreeNode{
			Val: 2,
		},
	}

	// Case 3: [1,null,2,null,3]
	rootThree := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Right: &TreeNode{
				Val: 3,
			},
		},
	}

	// Case 4: [1,2]
	rootFour := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}

	fmt.Println(isBalanced(root))
	fmt.Println(isBalanced(rootTwo))
	fmt.Println(isBalanced(rootThree))
	fmt.Println(isBalanced(rootFour))
}
