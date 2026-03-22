package main

func findRotation(mat [][]int, target [][]int) bool {
	n := len(mat)
	if isSame(mat, target) {
		return true
	}

	tmp := make([][]int, n)
	for i := range n {
		tmp[i] = make([]int, n)
	}
	for range 3 {
		rotateMat(mat, tmp, n)
		if isSame(mat, tmp) {
			return true
		}
		mat, tmp = tmp, mat
	}
	return false
}

func rotateMat(mat, out [][]int, n int) {
	for i := range n {
		for j := range n {
			out[i][j] = mat[n-j-1][i]
		}
	}
}

func isSame(a, b [][]int) bool {
	for i, row := range a {
		for j, x := range row {
			if x != b[i][j] {
				return false
			}
		}
	}
	return true
}

func findRotationV1(mat [][]int, target [][]int) bool {
	oks := make([]bool, 4)
	n := len(mat)
	for i := range oks {
		oks[i] = true
	}
	for i, row := range mat {
		for j, x := range row {
			oks[0] = oks[0] && (x == target[i][j])
			oks[1] = oks[1] && (x == target[n-j-1][i])
			oks[2] = oks[2] && (x == target[n-i-1][n-j-1])
			oks[3] = oks[3] && (x == target[j][n-i-1])
		}
	}
	for _, ok := range oks {
		if ok {
			return true
		}
	}
	return false
}
