package graph

import "container/heap"

// dijkstra 查找 start 到 end 节点中最短路长度
func dijkstra(g *Graph, start int, end int) int {
	dist := make([]int, g.n)
	for i := range g.n {
		dist[i] = 0x3f3f3f3f
	}
	dist[start] = 0
	pq := &hq{{start, 0}}
	for pq.Len() > 0 {
		p := heap.Pop(pq).(state)
		u, d := p.to, p.dis
		if dist[u] < d { // 已经找到更短的路径，跳过
			continue
		}
		if end == u { // 找到终点
			continue
		}
		// 从节点u出发，尝试找到一条比现有更短的路径
		for i := g.f[u]; i != -1; i = g.ne[i] {
			v := g.v[i]
			target := d + g.w[i]
			if dist[v] > target {
				dist[v] = target
				heap.Push(pq, state{v, target})
			}
		}
	}
	return -1
}

type (
	state struct{ to, dis int }
	hq    []state
)

func (hq *hq) Len() int               { return len(*hq) }
func (hq *hq) Less(i int, j int) bool { return (*hq)[i].dis < (*hq)[j].dis }
func (hq *hq) Swap(i int, j int)      { (*hq)[i], (*hq)[j] = (*hq)[j], (*hq)[i] }
func (hq *hq) Push(x any)             { (*hq) = append((*hq), x.(state)) }
func (hq *hq) Pop() any               { n := len(*hq); item := (*hq)[n-1]; (*hq) = (*hq)[:n-1]; return item }
