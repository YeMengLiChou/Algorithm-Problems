package main

import "fmt"

func matrixBlockSum(mat [][]int, k int) [][]int {
	n, m := len(mat), len(mat[0])
	sum := make([][]int, n+1)
	res := make([][]int, n)
	for i := range sum {
		sum[i] = make([]int, m+1)
	}
	for k := range n {
		for j := range m {
			sum[k+1][j+1] = sum[k+1][j] + mat[k][j]
		}
	}
	for k := range m {
		for i := range n {
			sum[i+1][k+1] += sum[i][k+1]
		}
	}
	for i := range n {
		fmt.Println(sum[i+1])
		res[i] = make([]int, m)
		for j := range m {
			sx, sy, ex, ey := max(0, i-k), max(0, j-k), min(n-1, i+k), min(m-1, j+k)
			res[i][j] = sum[ex+1][ey+1] - sum[ex+1][sy] - sum[sx][ey+1] + sum[sx][sy]
		}
	}
	return res
}

func main() {
	res := matrixBlockSum(
		[][]int{
			{1, 2, 3}, {4, 5, 6}, {7, 8, 9},
		}, 1)
	fmt.Println(res)
}
