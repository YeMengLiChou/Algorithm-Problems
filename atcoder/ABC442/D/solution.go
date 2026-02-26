package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
)

type Int = int

func solution() {
	N, Q := readI(), readI()
	A := make([]int, N)
	for i := range N {
		A[i] = readI()
	}
	st := NewSegTree(A)
	for Q > 0 {
		Q--
		action := readI()
		switch action {
		case 1:
			x := readI()
			a, b := st.Get(x-1), st.Get(x)
			st.Update(x-1, x-1, -a+b)
			st.Update(x, x, -b+a)
		case 2:
			l, r := readI(), readI()
			puts(st.Query(l-1, r-1), "\n")
		}
	}
}

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

func (st *SegTree) Get(idx int) int {
	return st.get(1, 0, st.N-1, idx)
}

func (st *SegTree) get(u, l, r, idx int) int {
	if l == r {
		return st.Nodes[u].Val
	}
	mid, lc, rc := position(u, l, r)
	st.pushdown(u, lc, rc, l, r)
	if idx <= mid {
		return st.get(lc, l, mid, idx)
	} else {
		return st.get(rc, mid+1, r, idx)
	}
}

var (
	file = os.Getenv("input")
	in   = bufio.NewReader(os.Stdin)
	out  = bufio.NewWriter(os.Stdout)
)

func main() {
	if len(file) > 0 {
		_, err := os.Stat(file)
		out.WriteString("(use file)\n")
		if err == nil {
			f, _ := os.Open(file)
			in = bufio.NewReader(f)
		}
	}
	defer out.Flush()
	T := 1
	// T = readI()
	for range T {
		solution()
	}
}

func readI() Int { // 快读
	var (
		x   Int
		neg bool
		c   byte
		err error
	)
	for c, err = in.ReadByte(); c < '0' || c > '9'; c, err = in.ReadByte() {
		if c == '-' {
			neg = true
		}
		if err == io.EOF {
			return 0
		}
	}
	for c >= '0' && c <= '9' {
		x = x*10 + Int(c-'0')
		c, _ = in.ReadByte()
	}
	if neg {
		return -x
	}
	return x
}

func puts(x ...any) { // 快写
	for _, val := range x {
		fmt.Fprint(out, val)
	}
}

func readS() string { // 快读字符串
	var strBuffer bytes.Buffer
	var c byte
	var err error
	for {
		c, err = in.ReadByte()
		if err != nil {
			if err == io.EOF {
				return ""
			}
			return ""
		}
		if c == ' ' || c == '\n' || c == '\r' || c == '\t' {
			continue
		}
		break
	}
	strBuffer.WriteByte(c)
	for {
		c, err = in.ReadByte()
		if err != nil {
			if err == io.EOF {
				break
			}
		}
		if c == ' ' || c == '\n' || c == '\r' || c == '\t' {
			in.UnreadByte()
			break
		}
		strBuffer.WriteByte(c)
	}
	return strBuffer.String()
}
