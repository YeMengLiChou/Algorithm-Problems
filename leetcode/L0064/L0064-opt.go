package main

import "math"

// minPathSum 优化版本
func minPathSum(grid [][]int) int {
	const INF = math.MaxInt >> 1
	n, m := len(grid), len(grid[0])

	dp := make([]int, m+1)
	for i := range m + 1 {
		dp[i] = INF
	}
	dp[1] = 0
	for i := range n {
		for j := 1; j <= m; j++ {
			dp[j] = min(dp[j], dp[j-1]) + grid[i][j-1]
		}
	}
	return dp[m]
}
