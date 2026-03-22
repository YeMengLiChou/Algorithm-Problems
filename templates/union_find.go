package templates

type UnionFind struct {
	p    []int
	cnts []int // 每个连通块的节点数量
	size int   // 连通块数量
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

// Union 合并两个节点所属连通块，true 表示合并前不是同一个连通块
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
