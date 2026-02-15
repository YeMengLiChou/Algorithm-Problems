package main

import (
	"math"
	"slices"
)

// minCostOpt1 优化版本1
func minCostOpt1(grid [][]int, k int) int {
	const INF = math.MaxInt32 >> 1
	n, m := len(grid), len(grid[0])
	size := n * m
	dp := make([][]int, n)
	items := make([]item, 0, size)
	for i := range n {
		dp[i] = make([]int, m)
		for j := range m {
			dp[i][j] = INF
			items = append(items, item{i, j})
		}
	}

	// sort with grid
	slices.SortFunc(items, func(a item, b item) int {
		return grid[a.i][a.j] - grid[b.i][b.j]
	})

	for range k + 1 {
		minCost := INF
		// 双指针，处理出从高值grid传送到低值grid所需的最小值
		for l, r := size-1, size-1; l >= 0; l-- {
			minCost = min(minCost, dp[items[l].i][items[l].j])
			if l-1 >= 0 && grid[items[l].i][items[l].j] == grid[items[l-1].i][items[l-1].j] {
				continue
			}
			for r >= l {
				dp[items[r].i][items[r].j] = minCost
				r--
			}
			r = l - 1
		}
		dp[0][0] = 0
		// 普通的转移
		for i := range n {
			for j := range m {
				x, y := i-1, j-1
				if x >= 0 {
					dp[i][j] = min(dp[i][j], dp[x][j]+grid[i][j])
				}
				if y >= 0 {
					dp[i][j] = min(dp[i][j], dp[i][y]+grid[i][j])
				}
			}
		}
	}
	return dp[n-1][m-1]
}

type item struct{ i, j int }
