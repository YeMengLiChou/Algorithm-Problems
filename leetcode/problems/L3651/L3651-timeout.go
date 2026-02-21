package main

import (
	"math"
	"slices"
	"sort"
)

// minCostTimeout 超时版本
func minCostTimeout(grid [][]int, kk int) int {
	n, m := len(grid), len(grid[0])
	dp := make([][][]int, n)
	_values := make(SortItems, n*m)
	values := &_values
	idx := 0
	for i := range n {
		dp[i] = make([][]int, m)
		for j := range m {
			dp[i][j] = make([]int, kk+1)
			for k := range kk + 1 {
				dp[i][j][k] = math.MaxInt32 >> 2
			}
			(*values)[idx] = &Item{i, j, grid[i][j]}
			idx++
		}
	}
	for i := range kk + 1 {
		dp[0][0][i] = 0
	}
	// sort with grid
	sort.Sort(values)

	for k := range kk + 1 {
		for i := range n {
			for j := range m {
				cost := grid[i][j]
				x, y := i-1, j-1
				if x >= 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[x][j][k]+cost)
				}
				if y >= 0 {
					dp[i][j][k] = max(dp[i][j][k], dp[i][y][k]+cost)
				}
			}
		}
		// 处理 k == 0，保证至少一个循环。
		if k == kk {
			break
		}
		for i := range n {
			for j := range m {
				// use k, find [x, y] satisify >= grid[i][j]
				leftPos, _ := slices.BinarySearchFunc(
					*values,
					&Item{0, 0, grid[i][j]},
					func(e *Item, target *Item) int {
						if e.value == target.value {
							return 0
						} else if e.value < target.value {
							return -1
						} else {
							return 1
						}
					},
				)

				for leftPos < len(*values) {
					x, y := (*values)[leftPos].i, (*values)[leftPos].j
					dp[i][j][k+1] = min(dp[i][j][k+1], dp[x][y][k])
					leftPos++
				}
			}
		}
	}

	return slices.Min(dp[n-1][m-1])
}

type (
	Item      struct{ i, j, value int }
	SortItems []*Item
)

func (items *SortItems) Len() int { return len(*items) }
func (items *SortItems) Less(i int, j int) bool {
	return (*items)[i].value < (*items)[j].value
}

func (items *SortItems) Swap(i int, j int) {
	(*items)[i], (*items)[j] = (*items)[j], (*items)[i]
}
