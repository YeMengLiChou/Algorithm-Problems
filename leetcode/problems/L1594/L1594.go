package main

import "math"

// const M = 1e9 + 7
// func maxProductPath(grid [][]int) int {
// 	m, n := len(grid), len(grid[0])
// 	ans := -1
// 	var dfs func(int, int, int)

// 	dfs = func(x, y int, cnt int) {
// 		if x == m-1 && y == n-1 {
// 			ans = max(ans, cnt)
// 			return
// 		}
// 		if x+1 < m {
// 			dfs(x+1, y, grid[x+1][y]*cnt)
// 		}
// 		if y+1 < n {
// 			dfs(x, y+1, grid[x][y+1]*cnt)
// 		}
// 	}
// 	dfs(0, 0, grid[0][0])
// 	if ans >= 0 {
// 		return ans % M
// 	} else {
// 		return -1
// 	}
// }

const M = 1e9 + 7

func maxProductPath(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	dp := make([][][2]int, m)
	for i := range dp {
		dp[i] = make([][2]int, n)
	}
	for i, row := range grid {
		for j, x := range row {
			if i == 0 && j == 0 {
				dp[i][j] = [2]int{x, x}
				continue
			}
			mx, mn := math.MinInt, math.MaxInt
			if i-1 >= 0 {
				mx = max(dp[i-1][j][0]*x, dp[i-1][j][1]*x)
				mn = min(dp[i-1][j][0]*x, dp[i-1][j][1]*x)
			}
			if j-1 >= 0 {
				mx = max(mx, dp[i][j-1][0]*x, dp[i][j-1][1]*x)
				mn = min(mn, dp[i][j-1][0]*x, dp[i][j-1][1]*x)
			}
			dp[i][j] = [2]int{mx, mn}
		}
	}
	ans := dp[m-1][n-1][0]
	if ans < 0 {
		return -1
	} else {
		return ans % M
	}
}
