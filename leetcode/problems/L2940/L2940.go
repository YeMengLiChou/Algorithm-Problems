package main

func leftmostBuildingQueries(heights []int, queries [][]int) []int {
	st := NewSegTree(heights)
	n := len(queries)
	res := make([]int, n)
	for i, x := range queries {
		l, r := x[0], x[1]
		if l > r {
			l, r = r, l
		}
		lh, rh := heights[l], heights[r]
		if l == r || lh < rh {
			res[i] = r
		} else {
			res[i] = st.Query(1, 0, st.N-1, r+1, st.N-1, max(heights[l], heights[r]))
		}
	}
	return res
}

type (
	Node struct {
		Val int
	}
	SegTree struct {
		N     int
		Nodes []Node
	}
)

func (st *SegTree) pushup(u, lc, rc int) {
	uNode, lNode, rNode := &st.Nodes[u], &st.Nodes[lc], &st.Nodes[rc]
	uNode.Val = max(lNode.Val, rNode.Val)
}

func NewSegTree(nums []int) *SegTree {
	n := len(nums)
	st := &SegTree{n, make([]Node, n<<2)}
	st.build(1, 0, st.N-1, nums)
	return st
}

func position(u int, l, r int) (mid, lc, rc int) {
	return l + ((r - l) >> 1), u << 1, u<<1 | 1
}

func (st *SegTree) build(u, l, r int, nums []int) {
	node := &st.Nodes[u]
	if l == r {
		node.Val = nums[l]
		return
	}
	mid, lc, rc := position(u, l, r)
	st.build(lc, l, mid, nums)
	st.build(rc, mid+1, r, nums)
	st.pushup(u, lc, rc)
}

func (st *SegTree) Query(u, l, r, ql, qr, x int) int {
	node := &st.Nodes[u]
	if node.Val <= x {
		return -1
	}
	if r < ql || l > qr {
		return -1
	}
	if l == r {
		return l
	}
	mid, lc, rc := position(u, l, r)
	val := st.Query(lc, l, mid, ql, qr, x)
	if val == -1 {
		val = st.Query(rc, mid+1, r, ql, qr, x)
	}
	return val
}
