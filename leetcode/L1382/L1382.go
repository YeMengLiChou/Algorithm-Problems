package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func balanceBST(root *TreeNode) *TreeNode {
	nums := make([]int, 0)
	dfs(root, &nums)
	return buildBalanceBST(nums, 0, len(nums))
}

func dfs(u *TreeNode, nums *[]int) {
	if u == nil {
		return
	}
	if u.Left != nil {
		dfs(u.Left, nums)
	}
	(*nums) = append(*nums, u.Val)
	if u.Right != nil {
		dfs(u.Right, nums)
	}
}

func buildBalanceBST(nums []int, l, r int) *TreeNode {
	if l >= r {
		return nil
	}
	mid := l + (r-l)>>1
	root := &TreeNode{nums[mid], nil, nil}
	root.Left = buildBalanceBST(nums, l, mid)
	root.Right = buildBalanceBST(nums, mid+1, r)
	return root
}
