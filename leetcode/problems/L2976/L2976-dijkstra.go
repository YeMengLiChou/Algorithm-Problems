package main

import (
	"container/heap"
	"math"
)

// minimumCost dijkstra 版本
func minimumCost(source string, target string, original []byte, changed []byte, cost []int) int64 {
	g := NewGraph(26, len(changed)+26)
	for i, from := range original {
		to, value := changed[i], cost[i]
		g.AddEdge(from-'a', to-'a', value)
	}
	allMinCost := make([][]int64, 0, 26)
	for i := range 26 {
		allMinCost = append(allMinCost, g.findMiniumnCost(byte(i)))
	}

	var ans int64
	for i := range source {
		from, to := source[i], target[i]
		if from == to {
			continue
		}
		value := allMinCost[from-'a'][to-'a']
		if value == -1 {
			return -1
		}
		ans += value
	}
	return ans
}

type Graph struct {
	n, idx     int
	f, w, next []int
	v          []byte
	dist       []int64
}

func NewGraph(n, m int) *Graph {
	f := make([]int, n+1)
	for i := range f {
		f[i] = -1
	}
	return &Graph{
		n:    n,
		idx:  0,
		f:    f,
		v:    make([]byte, m),
		w:    make([]int, m),
		next: make([]int, m),
		dist: make([]int64, n+1),
	}
}

func (g *Graph) AddEdge(u, v byte, w int) {
	idx := g.idx
	g.v[idx] = v
	g.w[idx] = w
	g.next[idx] = g.f[u]
	g.f[u] = idx
	g.idx++
}

func (g *Graph) findMiniumnCost(from byte) []int64 {
	const INF = math.MaxInt64 >> 1
	dist := make([]int64, g.n)
	for i := range dist {
		dist[i] = INF
	}
	dist[from] = 0
	pq := &hq{{from, 0}}

	for (*pq).Len() > 0 {
		p := heap.Pop(pq).(Item)
		u, dis := p.u, p.w
		if dist[u] > dis {
			continue
		}
		for i := g.f[u]; i != -1; i = g.next[i] {
			v := g.v[i]
			t := dis + int64(g.w[i])
			if dist[v] > t {
				dist[v] = t
				heap.Push(pq, Item{v, t})
			}
		}
	}
	for i, x := range dist {
		if x == INF {
			dist[i] = -1
		}
	}
	return dist
}

type (
	Item struct {
		u byte
		w int64
	}
	hq []Item
)

func (hq *hq) Len() int               { return len(*hq) }
func (hq *hq) Less(i int, j int) bool { return (*hq)[i].w < (*hq)[j].w }
func (hq *hq) Swap(i int, j int)      { (*hq)[i], (*hq)[j] = (*hq)[j], (*hq)[i] }
func (hq *hq) Push(x any)             { (*hq) = append((*hq), x.(Item)) }
func (hq *hq) Pop() any               { n := len(*hq); old := (*hq)[n-1]; (*hq) = (*hq)[:n-1]; return old }
