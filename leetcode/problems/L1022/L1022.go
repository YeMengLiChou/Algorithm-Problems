package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sumRootToLeaf(root *TreeNode) int {
	res := 0
	if root != nil {
		dfs(&res, 0, root)
	}
	return res
}

func dfs(sum *int, cnt int, root *TreeNode) {
	cnt = (cnt << 1) | root.Val
	if root.Left == nil && root.Right == nil {
		(*sum) += cnt
		return
	}
	if root.Left != nil {
		dfs(sum, cnt, root.Left)
	}
	if root.Right != nil {
		dfs(sum, cnt, root.Right)
	}
}
