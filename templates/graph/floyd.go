// Package graph
package graph

import "math"

func floyd(n int, g *Graph) [][]int {
	const INF = math.MaxInt >> 1
	dist := make([][]int, n+1)
	for i := range dist {
		dist[i] = make([]int, n+1)
		for j := range n + 1 {
			dist[i][j] = INF
		}
		dist[i][i] = 0
	}
	// 根据已有边初始化
	for i := 1; i <= n; i++ {
		for j := g.f[i]; j != -1; j = g.ne[j] {
			v := g.v[j]
			dist[i][v] = g.w[j]
		}
	}
	// 中间结点需要在最外层
	for k := 1; k <= n; k++ {
		for i := 1; i <= n; i++ {
			if dist[i][k] == INF {
				continue // 可以提前跳过不连通的边
			}
			for j := 1; j <= n; j++ {
				dist[i][j] = min(dist[i][j], dist[i][k]+dist[k][i])
			}
		}
	}
	return dist
}
