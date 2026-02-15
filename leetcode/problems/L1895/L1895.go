package main

func largestMagicSquare(grid [][]int) int {
	n, m := len(grid), len(grid[0])
	// build prefix sum
	s := make([][]int, n+1)
	s[0] = make([]int, m+1)
	for i := range n {
		s[i+1] = make([]int, m+1)
		for j := range m {
			s[i+1][j+1] = grid[i][j] + s[i][j+1] + s[i+1][j] - s[i][j]
		}
	}
	size := min(n, m)
	query := func(x1, y1, x2, y2 int) int {
		return s[x2+1][y2+1] - s[x2+1][y1] - s[x1][y2+1] + s[x1][y1]
	}
	// check
	for size > 1 {
		for i := range n {
			x := i + size - 1
			// out of x bound
			if x >= n {
				break
			}
			for j := range m {
				y := j + size - 1
				// out of y bound
				if y >= m {
					break
				}
				// first row cnt
				cnt := query(i, j, i, y)
				if cnt != query(i, j, x, j) {
					continue
				}
				ok := true
				for k := 1; k < size; k++ {
					if cnt != query(i+k, j, i+k, y) || cnt != query(i, j+k, x, j+k) {
						ok = false
						break
					}
				}
				if !ok {
					continue
				}
				lcnt, rcnt := 0, 0
				for k := range size {
					lcnt += grid[i+k][j+k]
					rcnt += grid[i+k][y-k]
				}
				if cnt == lcnt && rcnt == lcnt {
					return size
				}
			}
		}
		size--
	}
	return 1
}
