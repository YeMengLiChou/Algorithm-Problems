package main

func maxSubArray(nums []int) int {
	st := NewSegementTree(nums)
}

type (
	Node struct {
		L, R                          int // 区间
		LMaxSum, RMaxSum, MaxSum, Sum int
	}
	SegmentTree struct {
		Nodes []Node
		N     int
	}
)

func NewSegementTree(nums []int) *SegmentTree {
	n := len(nums)
	tree := &SegmentTree{make([]Node, n<<2), n}
	tree.build(1, 0, n-1, nums)
	return tree
}

func (st *SegmentTree) build(u, l, r int, nums []int) {
	if l == r {
		node := &st.Nodes[u]
		node.Val = nums[l]
		node.L = l
		node.R = r
		return
	}
	mid, lc, rc := l+((r-l)>>1), u<<1, u<<1|1
	st.build(lc, l, mid, nums)
	st.build(rc, mid+1, r, nums)
	st.pushup(u, lc, rc)
}

func (st *SegmentTree) Update(u, l, r, delta int) {
	node := &st.Nodes[u]
	if l <= node.L && node.R <= r {
		node.Val += (node.R - node.L + 1) * delta
		return
	}
	mid, lc, rc := node.L+((node.R-node.L)>>1), u<<1, u<<1|1
	st.pushdown(u, lc, rc)
	if l <= mid {
		st.Update(lc, l, mid, delta)
	}
	if r > mid {
		st.Update(rc, mid+1, r, delta)
	}
	st.pushup(u, lc, rc)
}

func (st *SegmentTree) Query(u, l, r int) int {
	node := &st.Nodes[u]
	if l <= node.L && node.R <= r {
		return node.Val
	}
	mid, lc, rc, val := node.L+((node.R-node.L)>>1), u<<1, u<<1|1, 0
	st.pushdown(u, lc, rc)
	if l <= mid {
		val += st.Query(lc, l, mid)
	}
	if r > mid {
		val += st.Query(rc, mid+1, r)
	}
	return val
}

func (st *SegmentTree) pushup(u, left, right int) {
	pNode, leftNode, rightNode := &st.Nodes[u], &st.Nodes[left], &st.Nodes[right]

	pNode.Val = leftNode.Val + rightNode.Val
	pNode.L = leftNode.L
	pNode.R = rightNode.R
}

func (st *SegmentTree) pushdown(p, left, right int) {
}
