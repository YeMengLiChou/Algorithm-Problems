package main

import (
	"math"
	"slices"
)

func minAbsDiff(grid [][]int, k int) [][]int {
	n, m := len(grid), len(grid[0])
	res := make([][]int, n-k+1)
	for i := range res {
		res[i] = make([]int, m-k+1)
	}
	idx, mx := 0, k*k
	subs := make([]int, mx)
	for i, row := range res {
		for j := range row {
			idx = 0
			for x := range k {
				for y := range k {
					subs[idx] = grid[i+x][j+y]
					idx++
				}
			}
			slices.Sort(subs)
			if isSame(subs) {
				row[j] = 0
			} else {
				mn := math.MaxInt
				for x := range mx - 1 {
					diff := subs[x+1] - subs[x]
					if diff == 0 {
						continue
					}
					mn = min(mn, diff)
				}
				row[j] = mn
			}
		}
	}
	return res
}

func isSame(x []int) bool {
	n := len(x)
	for i := range n - 1 {
		if x[i] != x[i+1] {
			return false
		}
	}
	return true
}
