package main

func countSubmatrices(grid [][]int, k int) int {
	n, m := len(grid), len(grid[0])
	sum := make([][]int, n+1)
	for i := range sum {
		sum[i] = make([]int, m+1)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			sum[i][j] = sum[i-1][j] + grid[i-1][j-1]
		}
	}
	res := 0
	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			sum[i][j] += sum[i][j-1]
			if sum[i][j] <= k {
				res++
			}
		}
	}

	return res
}

func countSubmatricesV1(grid [][]int, k int) int {
	cols := make([]int, len(grid[0]))
	res := 0
	for _, row := range grid {
		cnt := 0
		for j, x := range row {
			cols[j] += x
			cnt += cols[j]
			if cnt <= k {
				res++
			}
		}
	}
	return res
}
