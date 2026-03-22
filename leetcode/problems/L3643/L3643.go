package main

func reverseSubmatrix(grid [][]int, x int, y int, k int) [][]int {
	n := len(grid)
	end := y + k
	for i, j := x, n-1; i >= j; {
		for idx := y; idx <= end; idx++ {
			grid[i][idx], grid[j][idx] = grid[j][idx], grid[i][idx]
		}
	}
	return grid
}
