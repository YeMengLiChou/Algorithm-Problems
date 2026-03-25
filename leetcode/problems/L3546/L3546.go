package main

import "fmt"

func main() {
	fmt.Println(
		canPartitionGrid([][]int{
			{1, 3}, {2, 4},
		}),
	)
}

func canPartitionGrid(grid [][]int) bool {
	n, m := len(grid), len(grid[0])
	s := make([][]int, n+1)
	for i := range s {
		s[i] = make([]int, m+1)
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			s[i][j] = s[i][j-1] + grid[i-1][j-1]
		}
	}
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			s[i][j] += s[i-1][j]
		}
	}

	for i := 1; i <= n; i++ {
		if s[i][m] == s[n][m]-s[i][m] {
			return true
		}
	}
	for j := 1; j <= m; j++ {
		if s[n][j] == s[n][m]-s[n][j] {
			return true
		}
	}
	return false
}
