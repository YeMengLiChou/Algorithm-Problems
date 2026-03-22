package main

func maxStability(n int, edges [][]int, k int) int {
	mustUF := NewUnionFind(n)
	allUF := NewUnionFind(n)
	maxEdge, minEdge := 0, 1<<20
	for _, e := range edges {
		u, v, s, must := e[0], e[1], e[2], e[3]
		if must > 0 && !mustUF.Union(u, v) {
			// 必选边存在环
			return -1
		}
		allUF.Union(u, v)
		maxEdge = max(maxEdge, s)
		minEdge = min(minEdge, s)
	}
	if allUF.size > 1 {
		// 图不连通
		return -1
	}

	check := func(x int) bool {
		uf := NewUnionFind(n)
		for _, e := range edges {
			u, v, s, must := e[0], e[1], e[2], e[3]
			if must > 0 && s < x {
				// 必选边小于 x
				return false
			}
			if must > 0 || s >= x {
				uf.Union(u, v)
			}
		}
		kk := k
		for _, e := range edges {
			if kk == 0 || uf.size == 1 {
				break
			}
			u, v, s, must := e[0], e[1], e[2], e[3]
			if must == 0 && s < x && s*2 >= x && uf.Union(u, v) {
				kk--
			}
		}
		return uf.size == 1
	}

	// 二分
	left, right := minEdge, maxEdge*2+1
	for left < right {
		mid := (left + right) >> 1
		if check(mid) {
			left = mid
		} else {
			right = mid - 1
		}
	}
	return left
}

type UnionFind struct {
	p, cnts []int
	size    int // 连通块数量
}

func NewUnionFind(n int) *UnionFind {
	p := make([]int, n)
	cnt := make([]int, n)
	for i := range n {
		p[i] = i
		cnt[i] = 1
	}
	return &UnionFind{
		p, cnt, n,
	}
}

func (uf *UnionFind) Find(x int) int {
	if uf.p[x] != x {
		uf.p[x] = uf.Find(uf.p[x])
	}
	return uf.p[x]
}

func (uf *UnionFind) Union(x, y int) bool {
	px, py := uf.Find(x), uf.Find(y)
	if px == py {
		return false
	}
	// 保证 px 为节点数最多的集合
	if uf.cnts[px] < uf.cnts[py] {
		px, py = py, px
	}
	uf.p[py] = px
	uf.cnts[px] += uf.cnts[py]
	uf.size--
	return true
}
