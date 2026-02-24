package main

func longestBalancedV2(nums []int) int {
	last := map[int]int{}
	st := NewSegTree(make([]int, len(nums)+1))
	res, sum := 0, 0
	for i, x := range nums {
		i += 1
		v := 1
		if x&1 == 1 {
			v = -1
		}
		pre, ok := last[x]
		if !ok {
			sum += v
			// 第一次遇见，需要加上
			st.Update(1, 0, st.N-1, i, st.N-1, v)
		} else {
			// 再次遇见，撤销前一个位置
			st.Update(1, 0, st.N-1, pre, i-1, -v)
		}
		last[x] = i
		// 查询第一个等于 sum 的下标
		j := st.Query(1, 0, st.N-1, 0, i-1-res, sum)
		if j >= 0 {
			res = i - j
		}
	}
	return res
}

func longestBalanced(nums []int) int {
	return longestBalancedV2(nums)
}

type (
	Node struct {
		Lazy int
		Min  int
		Max  int
	}
	SegTree struct {
		N     int
		Nodes []Node
	}
)

func (st *SegTree) pushup(u, lc, rc int) {
	uNode, lNode, rNode := &st.Nodes[u], &st.Nodes[lc], &st.Nodes[rc]
	uNode.Max = max(lNode.Max, rNode.Max)
	uNode.Min = min(lNode.Min, rNode.Min)
}

func (st *SegTree) pushdown(u, lc, rc, l, r int) {
	uNode := &st.Nodes[u]
	lazy := uNode.Lazy
	if lazy == 0 {
		return
	}
	lNode, rNode := &st.Nodes[lc], &st.Nodes[rc]
	lNode.Lazy += lazy
	lNode.Max += lazy
	lNode.Min += lazy
	rNode.Lazy += lazy
	rNode.Max += lazy
	rNode.Min += lazy
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
		node.Max = nums[l]
		node.Min = nums[l]
		return
	}
	mid, lc, rc := position(u, l, r)
	st.build(lc, l, mid, nums)
	st.build(rc, mid+1, r, nums)
	st.pushup(u, lc, rc)
}

func (st *SegTree) Update(u, l, r, ql, qr, k int) {
	if r < ql || l > qr {
		return
	} else if ql <= l && r <= qr {
		node := &st.Nodes[u]
		node.Lazy += k
		node.Max += k
		node.Min += k
		return
	}
	mid, lc, rc := position(u, l, r)
	st.pushdown(u, lc, rc, l, r)
	st.Update(lc, l, mid, ql, qr, k)
	st.Update(rc, mid+1, r, ql, qr, k)
	st.pushup(u, lc, rc)
}

// Query 找到第一个等于 x 的下标
func (st *SegTree) Query(u, l, r, ql, qr, x int) int {
	if r < ql || l > qr || (x < st.Nodes[u].Min || x > st.Nodes[u].Max) {
		return -1
	}
	if l == r {
		return l
	}
	mid, lc, rc := position(u, l, r)
	st.pushdown(u, lc, rc, l, r)
	if idx := st.Query(lc, l, mid, ql, qr, x); idx >= 0 {
		return idx
	}
	return st.Query(rc, mid+1, r, ql, qr, x)
}
