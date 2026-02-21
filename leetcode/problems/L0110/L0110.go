package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isBalanced(root *TreeNode) bool {
	_, ok := getHeight(root)
	return ok
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return -x
}

func getHeight(root *TreeNode) (int, bool) {
	if root == nil {
		return 0, true
	}
	lh, ok := getHeight(root.Left)
	if !ok {
		return -1, false
	}
	rh, ok := getHeight(root.Right)
	if !ok {
		return -1, false
	}
	if abs(lh-rh) > 1 {
		return -1, false
	}
	return max(lh, rh) + 1, true
}
