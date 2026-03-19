package main

func numberOfSubmatrices(grid [][]byte) int {
	n, m := len(grid), len(grid[0])
	colX, colY := make([]int, m), make([]int, m)
	res := 0
	for i := range n {
		xCnt, yCnt := 0, 0
		for j := range m {
			switch grid[i][j] {
			case 'X':
				colX[j] += 1
			case 'Y':
				colY[j] += 1
			}
			xCnt += colX[j]
			yCnt += colY[j]
			if xCnt >= 1 && xCnt == yCnt {
				res++
			}
		}
	}
	return res
}
