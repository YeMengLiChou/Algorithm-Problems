package main

func numSpecial(mat [][]int) int {
	m, n := len(mat), len(mat[0])
	cols := make([]int, n)
	for i := range n {
		for j := range m {
			cols[i] += mat[j][i]
		}
	}
	res := 0
	for i := range n {
		hasOne := false
		for j := range m {
			if mat[i][j] == 1 {
				if hasOne {
					res--
					break
				}
				if cols[j] > 1 {
					break
				} else {
					res++
					hasOne = true
				}
			}
		}
	}
	return res
}
