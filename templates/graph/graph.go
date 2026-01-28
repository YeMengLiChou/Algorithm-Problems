package graph

const INF = 0x3f3f3f3f

// =============== 二维矩阵 ============
// type Graph struct {
// 	n  int
// 	mp [][]int
// }

// func NewGraph(n int) *Graph {
// 	mp := make([][]int, n)
// 	for i := range mp {
// 		elem := make([]int, n)
// 		for j := range elem {
// 			elem[j] = INF
// 		}
// 		mp[i] = elem
// 	}
// 	return &Graph{n, mp}
// }

// =============== 邻接表 ==============

// type (
// 	Edge  struct{ v, w int }
// 	Graph struct {
// 		n   int
// 		adj [][]Edge
// 		vis []bool
// 	}
// )

// func NewGraph(n int) *Graph {
// 	adj := make([][]Edge, n)
// 	vis := make([]bool, n)
// 	return &Graph{n, adj, vis}
// }

// func (g *Graph) addEdge(u, v, w int) {
// 	g.adj[u] = append(g.adj[u], Edge{v, w})
// }

// func (g *Graph) findEdge(u, v int) bool {
// 	return slices.IndexFunc(g.adj[u], func(e Edge) bool {
// 		return e.v == v
// 	}) >= 0
// }

// func (g *Graph) dfs(u int) {
// 	if g.vis[u] {
// 		return
// 	}
// 	g.vis[u] = true
// 	for _, v := range g.adj[u] {
// 		g.dfs(v.v)
// 	}
// }

// =============== 链式向前星 ==============

type Graph struct {
	n, idx      int
	f, v, w, ne []int
	vis         []bool
}

func NewGraph(n, m int) *Graph {
	f := make([]int, n+1)
	for i := range f {
		f[i] = -1
	}
	return &Graph{
		n:   n + 1,
		f:   f,
		v:   make([]int, m),
		w:   make([]int, m),
		ne:  make([]int, m),
		vis: make([]bool, n),
	}
}

func (g *Graph) addEdge(u, v, w int) {
	idx := g.idx
	g.v[idx] = v
	g.w[idx] = w
	g.ne[idx] = g.f[u]
	g.f[u] = idx
	g.idx++
}
