package main

func minSwaps(grid [][]int) int {
	n := len(grid)
	tails := make([]int, n)
	for i := range n {
		j := n - 1
		for j >= 0 && grid[i][j] == 0 {
			j--
		}
		if j == -1 {
			return -1
		}
		tails[i] = j
	}
	res := 0
	for i := range n {
		k := -1
		for j := i; j < n; j++ {
			if tails[j] <= i {
				res += j - i
				k = j
				break
			}
		}
		if k == -1 {
			return -1
		} else {
			for j := k; j > i; j-- {
				tails[j-1], tails[j] = tails[j], tails[j-1]
			}
		}
	}
	return res
}
