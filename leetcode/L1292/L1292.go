package main

func maxSideLength(mat [][]int, threshold int) int {
	n, m := len(mat), len(mat[0])
	s := make([][]int, n+1)

	// build prefix sum
	s[0] = make([]int, m+1)
	for i := range n {
		s[i+1] = make([]int, m+1)
		for j := range m {
			s[i+1][j+1] = mat[i][j] + s[i][j+1] + s[i+1][j] - s[i][j]
		}

	}

	check := func(size int) bool {
		for i := range n {
			for j := range m {
				bx, by := i+size, j+size
				if bx > n || by > m {
					continue
				}
				sum := s[bx][by] - s[bx][j] - s[i][by] + s[i][j]
				if sum <= threshold {
					return true
				}
			}
		}
		return false
	}

	// bineary search
	l, r, mid := 0, min(n, m), 0
	for l <= r {
		mid = l + (r-l)>>1
		if check(mid) {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return l - 1
}
