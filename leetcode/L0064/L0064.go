package main

import "math"

func minPathSum(grid [][]int) int {
	const INF = math.MaxInt >> 1
	n, m := len(grid), len(grid[0])

	dp := make([][]int, n+1)
	for i := range n + 1 {
		dp[i] = make([]int, m+1)
		for j := range m + 1 {
			dp[i][j] = INF
		}
	}
	dp[1][1] = grid[0][0]
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			dp[i][j] = min(
				dp[i][j],
				min(dp[i-1][j], dp[i][j-1])+grid[i-1][j-1],
			)
		}
	}
	return dp[n][m]
}
