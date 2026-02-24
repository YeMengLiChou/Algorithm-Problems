package templates

type (
	Node struct {
		Lazy int
		Val  int
	}
	SegTree struct {
		N     int
		Nodes []Node
	}
)

func (st *SegTree) pushup(u, lc, rc int) {
	uNode, lNode, rNode := &st.Nodes[u], &st.Nodes[lc], &st.Nodes[rc]
	uNode.Val = lNode.Val + rNode.Val
}

func (st *SegTree) pushdown(u, lc, rc, l, r int) {
	uNode := &st.Nodes[u]
	lazy := uNode.Lazy
	if lazy == 0 {
		return
	}
	mid, lNode, rNode := l+((r-l)>>1), &st.Nodes[lc], &st.Nodes[rc]
	lNode.Lazy += lazy
	lNode.Val += (mid - l + 1) * lazy
	rNode.Lazy += lazy
	rNode.Val += (r - mid) * lazy
	uNode.Lazy = 0
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

func (st *SegTree) Update(ql, qr, k int) {
	st.update(1, 0, st.N-1, ql, qr, k)
}

func (st *SegTree) update(u, l, r, ql, qr, k int) {
	if r < ql || l > qr {
		return
	} else if ql <= l && r <= qr {
		node := &st.Nodes[u]
		node.Lazy += k
		node.Val += (r - l + 1) * k
		return
	}
	mid, lc, rc := position(u, l, r)
	st.pushdown(u, lc, rc, l, r)
	st.update(lc, l, mid, ql, qr, k)
	st.update(rc, mid+1, r, ql, qr, k)
	st.pushup(u, lc, rc)
}

func (st *SegTree) Query(ql, qr int) int {
	return st.query(1, 0, st.N-1, ql, qr)
}

func (st *SegTree) query(u, l, r, ql, qr int) int {
	if r < ql || l > qr {
		return 0
	} else if ql <= l && r <= qr {
		return st.Nodes[u].Val
	}
	mid, lc, rc := position(u, l, r)
	st.pushdown(u, lc, rc, l, r)
	return st.query(lc, l, mid, ql, qr) + st.query(rc, mid+1, r, ql, qr)
}
