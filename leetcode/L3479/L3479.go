package main

func numOfUnplacedFruits(fruits []int, baskets []int) int {
	st := NewSegTree(baskets)
	res := 0
	for _, x := range fruits {
		idx := st.QueryAndUpdate(1, 0, st.N-1, x)
		if idx == -1 {
			res++
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

func (st *SegTree) QueryAndUpdate(u, l, r, mx int) int {
	node := &st.Nodes[u]
	if node.Val < mx {
		return -1
	}
	if l == r {
		node.Val = -1
		return l
	}
	mid, lc, rc := position(u, l, r)
	val := st.QueryAndUpdate(lc, l, mid, mx)
	if val == -1 {
		val = st.QueryAndUpdate(rc, mid+1, r, mx)
	}
	st.pushup(u, lc, rc)
	return val
}
