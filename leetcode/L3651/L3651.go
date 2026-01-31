package main

import (
	"math"
	"slices"
)

// minCost 优化版本2
func minCost(grid [][]int, k int) int {
	const INF = math.MaxInt32 >> 1
	n, m := len(grid), len(grid[0])
	dp := make([]int, m+1)
	mx := 0
	for i := range n {
		mx = max(mx, slices.Max(grid[i]))
	}
	sufMin := make([]int, mx+2)
	tmpMin := make([]int, mx+1)
	for i := range sufMin {
		sufMin[i] = INF
	}

	for range k + 1 {
		for i := range tmpMin {
			tmpMin[i] = INF
		}
		for i := range dp {
			dp[i] = INF
		}
		dp[1] = -grid[0][0] // 避免额外处理，在下面计算时会变成0
		for i := range n {
			for j, x := range grid[i] {
				// 普通转移+传送转移
				dp[j+1] = min(dp[j+1]+x, dp[j]+x, sufMin[x])
				// 不能影响到 sufMin，需要另外存
				tmpMin[x] = min(tmpMin[x], dp[j+1])
			}
		}
		// 后缀最小值
		for i := mx; i >= 0; i-- {
			sufMin[i] = min(sufMin[i+1], tmpMin[i])
		}
	}
	return dp[m]
}
