package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepthDfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	lDepth, rDepeth := minDepth(root.Left), minDepth(root.Right)
	if lDepth == 0 || rDepeth == 0 {
		return max(lDepth, rDepeth) + 1
	}
	return min(lDepth, rDepeth) + 1
}

func minDepthBfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	q := []*TreeNode{root}
	idx := 0
	depth := 0
	for len(q) > 0 {
		n := len(q) - idx
		depth += 1
		for n > 0 {
			u := q[idx]
			if u.Left == nil && u.Right == nil {
				return depth
			}
			if u.Left != nil {
				q = append(q, u.Left)
			}
			if u.Right != nil {
				q = append(q, u.Right)
			}
			idx++
			n--
		}
	}
	return depth
}

func minDepth(root *TreeNode) int {
	return minDepthBfs(root)
}
