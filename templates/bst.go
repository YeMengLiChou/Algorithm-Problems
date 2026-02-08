package templates

type BSTNode struct {
	key   int
	left  *BSTNode
	right *BSTNode
	// 维护其他信息，如高度，节点数量
	size  int // 当前节点为根的子节点数量
	count int // 当前节点的重复数量
}

func NewBSTNode(value int) *BSTNode {
	return &BSTNode{value, nil, nil, 1, 1}
}

func inorderTraversal(root *BSTNode, result *[]int) {
	if root == nil {
		return
	}
	inorderTraversal(root.left, result)
	(*result) = append((*result), root.key)
	inorderTraversal(root.right, result)
}
