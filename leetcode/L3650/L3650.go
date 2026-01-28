package main

import (
	"container/heap"
)

func minCost(n int, edges [][]int) int {
	g := NewGraph(n, len(edges)<<1)
	for _, e := range edges {
		u, v, w := e[0], e[1], e[2]
		g.addEdge(u, v, w)
		g.addEdge(v, u, w<<1) // 添加反转边

	}
	dist := make([]int, n)
	const INF = 0x3f3f3f3f
	for i := range dist {
		dist[i] = INF
	}
	dist[0] = 0

	pq := &hq{}
	heap.Push(pq, pair{0, 0})

	for pq.Len() > 0 {
		top := heap.Pop(pq).(pair)
		u, dis := top.to, top.dis
		if dis > dist[u] {
			continue
		}
		if u == n-1 {
			return dis
		}
		for i := g.f[u]; i != -1; i = g.ne[i] {
			v := g.v[i]
			d := g.w[i] + dis
			if dist[v] > d {
				dist[v] = d
				heap.Push(pq, pair{v, d})
			}
		}
	}
	return -1
}

type (
	pair struct{ to, dis int }
	hq   []pair
)

func (hq *hq) Len() int               { return len(*hq) }
func (hq *hq) Less(i int, j int) bool { return (*hq)[i].dis < (*hq)[j].dis }
func (hq *hq) Swap(i int, j int)      { (*hq)[i], (*hq)[j] = (*hq)[j], (*hq)[i] }
func (hq *hq) Push(x any)             { (*hq) = append((*hq), x.(pair)) }
func (hq *hq) Pop() any {
	n := len(*hq)
	item := (*hq)[n-1]
	(*hq) = (*hq)[:n-1]
	return item
}

type Graph struct {
	n, idx      int
	f, v, w, ne []int
}

func NewGraph(n, m int) *Graph {
	f := make([]int, n+1)
	for i := range f {
		f[i] = -1
	}
	return &Graph{
		n:  n + 1,
		f:  f,
		v:  make([]int, m),
		w:  make([]int, m),
		ne: make([]int, m),
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
